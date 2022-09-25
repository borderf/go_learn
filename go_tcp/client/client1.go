package main

import (
	"fmt"
	"go_learn/go_tcp/common"
	"net"
)

func main() {
	// 连接地址
	serverAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:5656")
	common.CheckError(err)
	conn, err := net.DialTCP("tcp4", nil, serverAddr)
	common.CheckError(err)
	fmt.Printf("connet to server %s\n", conn.RemoteAddr().String())

	conn.Write([]byte("hello|"))
	conn.Write([]byte("golang|"))

	conn.Close()
}
