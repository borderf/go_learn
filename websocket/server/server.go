package main

import (
	"fmt"
	"go_learn/go_tcp/common"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type WsServer struct {
	// listener net.Listener
	addr    string
	upgrade *websocket.Upgrader
}

func NewWsServer(port int) *WsServer {
	ws := new(WsServer)
	ws.addr = "0.0.0.0:" + strconv.Itoa(port)
	ws.upgrade = &websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
	}
	// 没给listener赋值
	return ws
}

func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 将http协议升级到websocket协议
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("升级失败 %s\n", err.Error())
		return
	}
	fmt.Printf("跟客户端 %s 建立好了 websocket 连接\n", r.RemoteAddr)
	go ws.handleOneConnection(conn)
}

func (ws *WsServer) handleOneConnection(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	for { // 长连接
		// 设置读deadline
		conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		var request common.Request
		if err := conn.ReadJSON(&request); err != nil {
			if netError, ok := err.(net.Error); ok {
				if netError.Timeout() {
					fmt.Println("发生了读超时")
					return
				}
			}
			fmt.Println("conn read err: ", err)
			return
		}
		response := common.Response{Sum: request.A + request.B}
		if err := conn.WriteJSON(response); err != nil {
			fmt.Println("conn send err: ", err)
		}

	}
}

func (ws *WsServer) Start() error {
	var err error
	// ws.listener, err = net.Listen("tcp", ws.addr)
	// if err != nil {
	// 	fmt.Printf("listen 失败 %s\n", err.Error())
	// 	return err
	// }
	// if err = http.Serve(ws.listener, ws); err != nil {
	// 	fmt.Printf("server 失败 %s\n", err.Error())
	// 	return err
	// }
	// 上面代码等价于listenAndServe代码
	if err = http.ListenAndServe(ws.addr, ws); err != nil {
		fmt.Printf("listen and server 失败 %s\n", err.Error())
		return err
	} else {
		return nil
	}
}

func main() {
	ws := NewWsServer(5657)
	ws.Start()
}
