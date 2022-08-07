package main

import (
	"fmt"
	"sync"
)

func main() {
	testMap()
}

// 会报错，原生map不支持并发写
func testMap() {
	dict := make(map[int]int)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		fmt.Println("thread is", i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dict[i] = i
		}(i)
	}
	wg.Wait()
	fmt.Println("over.......")
}
