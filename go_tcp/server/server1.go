package main

import (
	"fmt"
	"go_learn/go_tcp/common"
	"net"
	"strings"
)

func main() {
	// 连接地址
	serverAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:5656")
	common.CheckError(err)
	// 建立监听器
	listener, err := net.ListenTCP("tcp4", serverAddr)
	common.CheckError(err)
	// 阻塞建立连接
	conn, err := listener.Accept()
	common.CheckError(err)
	fmt.Printf("client %s connect to me\n", conn.RemoteAddr().String())

	// 获取请求内容
	request := make([]byte, 256)
	n, err := conn.Read(request)
	common.CheckError(err)

	arr := strings.Split(string(request[:n]), "|")
	for _, v := range arr {
		fmt.Println("value is", v)
	}
}

func readDataGram(conn net.Conn) []string {
	dataGram := make([]string, 0)
	content := make([]byte, 1024)
	n, err := conn.Read(content)
	common.CheckError(err)
	var j = 0
	beginIndex := 0
	for i := 0; i < n; i++ {
		if content[i] == common.Delemeter[j] {
			if j == len(common.Delemeter) {
				// 从content中完整匹配到delemeter
				// 获取到分隔符之前的内容
				dataGram = append(dataGram, string(content[beginIndex:i-len(common.Delemeter)]))
				beginIndex = i + 1
			}
			j++
		} else {
			j = 0
		}
	}
	return dataGram
}
