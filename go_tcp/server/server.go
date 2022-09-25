package main

import (
	"encoding/json"
	"fmt"
	"go_learn/go_tcp/common"
	"net"
)

func main1() {
	// 1、建立连接
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:5656")
	common.CheckError(err)
	// 监听
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	common.CheckError(err)
	// 并发处理多个连接
	for {
		// 建立连接，阻塞
		conn, err := listener.Accept()
		common.CheckError(err)
		go handleOneClient(conn)
	}

}

func handleOneClient(conn net.Conn) {
	defer conn.Close()
	for {
		// 获取请求数据
		request := make([]byte, 256)
		// 成功读到多少字节
		n, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		// 反序列化
		var r common.Request
		err = json.Unmarshal(request[:n], &r)
		common.CheckError(err)

		fmt.Printf("request A is %d, B is %d\n", r.A, r.B)

		// 写回响应
		response := common.Response{
			Sum: r.A + r.B,
		}
		bs, err := json.Marshal(response)
		common.CheckError(err)
		_, err = conn.Write(bs)
		common.CheckError(err)
	}

}
