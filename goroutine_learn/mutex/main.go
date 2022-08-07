package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex
	GOTIME := "2006-01-02 15:04:05"

	go func() {
		defer lock.Unlock()
		lock.Lock()
		fmt.Println("func1 get lock at", time.Now().Format(GOTIME))
		time.Sleep(time.Second)
		fmt.Println("func1 release lock at", time.Now().Format(GOTIME))
	}()
	time.Sleep(100 * time.Millisecond)
	go func() {
		defer lock.Unlock()
		lock.Lock()
		fmt.Println("func2 get lock at", time.Now().Format(GOTIME))
		time.Sleep(time.Second)
		fmt.Println("func2 release lock at", time.Now().Format(GOTIME))
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("over..........")
}
