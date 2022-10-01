package main

import (
	"fmt"
	"go_learn/go_tcp/common"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	// 拨号器
	dialer := websocket.Dialer{}
	header := http.Header{
		"Name": []string{"hjy"},
	}
	// 建立连接
	conn, resp, err := dialer.Dial("ws://localhost:5657", header)
	if err != nil {
		fmt.Printf("dial失败 %s\n", err.Error())
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	msg, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(msg))
	fmt.Println("response header")
	for k, v := range resp.Header {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}
	// 循环调用5次
	defer conn.Close()
	for i := 0; i < 5; i++ {
		request := common.Request{A: 3, B: 5}
		conn.WriteJSON(request)
		var response common.Response
		conn.ReadJSON(&response)
		fmt.Println(response.Sum)
	}

}
