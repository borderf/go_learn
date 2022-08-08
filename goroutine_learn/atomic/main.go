package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// testAdd()
	// testCompareAndSwap()
	// testLoad()
	// testStore()
	testSwap()
}

func testAdd() {
	var num int64
	fmt.Println("num is", num)
	fmt.Println(">>>>> add 10 is", atomic.AddInt64(&num, 10))
	fmt.Println("===== sub 2 is", atomic.AddInt64(&num, -2))
}

func testCompareAndSwap() {
	var num int64
	num = 10
	fmt.Println("num is", num)
	// old != new
	atomic.CompareAndSwapInt64(&num, 20, 12)
	fmt.Println("cas num is", num)
	// old == new
	atomic.CompareAndSwapInt64(&num, 10, 100)
	fmt.Println("cas num is", num)
}

// 原子性读取
func testLoad() {
	var num int64
	num = 10
	fmt.Println("num is", num)
	fmt.Println(atomic.LoadInt64(&num))
}

func testStore() {
	var num int64
	atomic.StoreInt64(&num, 100)
	fmt.Println(num)
}

func testSwap() {
	var num int64
	num = 10
	i := atomic.SwapInt64(&num, 100)
	fmt.Println("old value is", i)
	fmt.Println("num is", num)
}
