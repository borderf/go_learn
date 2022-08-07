package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.RWMutex
	GOTIME := "2006-01-02 15:04:05"

	// read lock
	fmt.Println("start read lock at", time.Now().Format(GOTIME))
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer lock.RUnlock()
			lock.RLock()
			fmt.Println("func", i, "get read lock at", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			fmt.Println("func", i, "release read lock at", time.Now().Format(GOTIME))
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// write lock
	fmt.Println("start write lock at", time.Now().Format(GOTIME))
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer lock.Unlock()
			lock.Lock()
			fmt.Println("func", i, "get write lock at", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			fmt.Println("func", i, "release write lock at", time.Now().Format(GOTIME))
		}(i)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("over..........")
}
