package main

import (
	"fmt"
	"time"
)

const GO_TIME_FORMAT = "2006-01-02 15:04:05"

func main() {
	testCh()
}

func testCh() {
	ch := make(chan string)
	go test1(ch)
	for {
		value, ok := <-ch
		if ok {
			fmt.Println("value is", value, "ok is", ok)
		} else {
			fmt.Println("channel is close")
			break
		}
	}
	fmt.Println("over......")

}

func test1(ch chan string) {
	defer func() {
		fmt.Println("i am test1, now close channel...")
		close(ch)
	}()
	for i := 0; i < 5; i++ {
		ch <- time.Now().Format(GO_TIME_FORMAT)
		time.Sleep(time.Second)
	}
}

func testSelect() {
	ch := make(chan string)
	go test1(ch)
	for {
		select {
		case value, ok := <-ch:
			if ok {
				fmt.Println("value is", value, "ok is", ok)
			} else {
				fmt.Println("ok is", ok)
				break
			}
		case <-time.After(2 * time.Second):
			fmt.Println("time out")
			break
		}
	}
}
