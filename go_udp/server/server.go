package main

import (
	"fmt"
	"go_learn/go_tcp/common"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5656")
	common.CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr) // 立刻返回一个【虚拟的】conn
	common.CheckError(err)

	defer conn.Close()

	for {
		requestBytes := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(requestBytes)
		if err != nil {
			fmt.Println("err is", err)
			// break
		}
		fmt.Printf("read %s from client %s \n", string(requestBytes[:n]), remoteAddr.String())
		// 返回原样的数据
		conn.WriteToUDP(requestBytes, remoteAddr)
	}

}
