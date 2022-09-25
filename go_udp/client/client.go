package main

import (
	"fmt"
	"go_learn/go_tcp/common"
	"net"
)

func main() {
	remoteAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5656")
	common.CheckError(err)
	conn, err := net.DialUDP("udp", nil, remoteAddr) // 立即成功返回，不需要阻塞
	common.CheckError(err)
	_, err = conn.Write([]byte("hello"))
	common.CheckError(err)

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	common.CheckError(err)
	fmt.Printf("get response %s\n", string(response[:n]))
	defer conn.Close()

}
