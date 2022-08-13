package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// get()
	complexRequest()
}

func get() {
	r, err := http.Get("http://127.0.0.1:8088/boy")
	if err != nil {
		panic(err)
	}
	// 关闭输出流，一定要关闭，否则协程泄露
	defer r.Body.Close()
	// 重定向到输出流
	io.Copy(os.Stdout, r.Body)
}

func complexRequest() {
	reader := strings.NewReader("hello server")
	if r, err := http.NewRequest("POST", "http://localhost:8088/boy", reader); err != nil {
		panic(err)
	} else {
		r.Header.Add("User-Agent", "lala")
		r.Header.Add("My-Header-Key", "MyHeader")
		// 自定义cokkie
		r.AddCookie(&http.Cookie{
			Name:    "auth",
			Value:   "password",
			Path:    "/",
			Domain:  "localhost",
			Expires: time.Now().Add(time.Duration(time.Hour)),
		})
		client := &http.Client{
			Timeout: 100 * time.Millisecond,
		}
		// 提交http请求
		res, err := client.Do(r)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			defer res.Body.Close()
			io.Copy(os.Stdout, res.Body)
		}
	}
}
