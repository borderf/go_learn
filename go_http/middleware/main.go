package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var limitChan = make(chan struct{}, 100)

type middleware func(http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		middlewareChain: make([]middleware, 0, 10),
		mux:             make(map[string]http.Handler, 10),
	}
}

// 添加中间件
func (router *Router) Use(m middleware) {
	router.middlewareChain = append(router.middlewareChain, m)
}

// 自定义路由
func (router *Router) Add(path string, handler http.Handler) {
	var mergeHandler = handler
	for i := 0; i < len(router.middlewareChain); i++ {
		mergeHandler = router.middlewareChain[i](mergeHandler)
	}
	router.mux[path] = mergeHandler
}

func main() {
	router := NewRouter()
	router.Use(timeMiddleWare)
	router.Use(limitMiddleware)

	router.Add("/", http.HandlerFunc(get))
	for path, handler := range router.mux {
		http.Handle(path, handler)
	}
	http.ListenAndServe(":8088", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	w.Write([]byte("how are you"))
}

// 记录耗时
func timeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		fmt.Println("begin is", begin)
		next.ServeHTTP(w, r)
		timeElapsed := time.Since(begin)
		log.Printf("request %s use %d ms \n", r.URL.Path, timeElapsed.Milliseconds())
	})
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limitChan <- struct{}{}
		next.ServeHTTP(w, r)
		<-limitChan
	})
}
