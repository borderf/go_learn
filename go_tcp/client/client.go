package main

import (
	"encoding/json"
	"fmt"
	"go_learn/go_tcp/common"
	"net"
)

func main1() {
	// 拨号建立连接
	serverAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:5656")
	common.CheckError(err)
	conn, err := net.DialTCP("tcp", nil, serverAddr)
	common.CheckError(err)
	for i := 0; i < 3; i++ {
		// 发送数据
		request := common.Request{
			A: 3,
			B: 5,
		}
		bs, err := json.Marshal(request)
		common.CheckError(err)
		_, err = conn.Write(bs)
		common.CheckError(err)
		// 接收数据
		response := make([]byte, 256)
		n, err := conn.Read(response)
		common.CheckError(err)

		var res common.Response
		err = json.Unmarshal(response[:n], &res)
		common.CheckError(err)
		fmt.Printf("response sum is %d\n", res.Sum)
	}
	// 关闭连接
	conn.Close()
}
