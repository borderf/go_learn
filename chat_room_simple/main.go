package main

import (
	"flag"
	"fmt"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	// 请求根目录时直接返回html页面
	http.ServeFile(w, r, "chat_room_simple/home.html")
}

func main() {
	// 命令行不指定port参数，则默认5657
	port := flag.String("port", "5657", "http service port")
	flag.Parse()
	hub := NewHub()
	go hub.Run()
	// 定义路由，注册每种请求对应的处理函数
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
	fmt.Printf("http serve on port %s\n", *port)
	// 启动成功，该行会一直阻塞
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		fmt.Printf("start http service error: %s\n", err)
	}
}
