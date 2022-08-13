package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", get)
	router.POST("/", post)

	http.ListenAndServe(":8088", router)
}

func handler(method string, w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("request method is", r.Method)
	fmt.Println("request body is")
	io.Copy(os.Stdout, r.Body)
	w.Write([]byte("hello boys, your request method is " + r.Method))
}

func get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	handler("GET", w, r, params)
}

func post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	handler("POST", w, r, params)
}
