package main

import (
	"fmt"
	"sync"
	"time"
)

const GO_TIME_FORMAT = "2006-01-02 15:04:05"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("func", i, "start at", time.Now().Format(GO_TIME_FORMAT))
			time.Sleep(time.Second)
			fmt.Println("func", i, "end at", time.Now().Format(GO_TIME_FORMAT))
		}(i)
	}
	wg.Wait()
	fmt.Println("over...........")
}
