package main

import (
	"fmt"
	"sync"
)

func main() {
	// testMap()
	// testSyncMap()
	sMap()
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

func testSyncMap() {
	var dict sync.Map
	var wg sync.WaitGroup

	// store
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dict.Store(i, fmt.Sprintf("%daaa", i))
		}(i)
	}
	wg.Wait()

	// load
	value, ok := dict.Load(1)
	fmt.Println("value is", value, "ok is", ok)

	// load or store
	actual, loaded := dict.LoadOrStore(200, "200aaa")
	fmt.Println("actual is", actual, "loaded is", loaded)

	actual2, loaded2 := dict.LoadOrStore(200, "200aaa")
	fmt.Println("actual2 is", actual2, "loaded2 is", loaded2)

	// delete
	dict.Delete(200)
	dict.Delete(20000)

	// load and delete
	value2, loaded3 := dict.LoadAndDelete(1)
	fmt.Println("load and delete value2 is", value2, "loaded3 is", loaded3)

	value3, loaded4 := dict.LoadAndDelete(1000)
	fmt.Println("load and delete value3 is", value3, "loaded4 is", loaded4)

	fmt.Println("over.....")
}

type Person struct {
	name string
	age  int
}

func sMap() {
	var dict sync.Map
	// string
	dict.Store("aaa", "AAA")
	// int
	dict.Store(1, 1)
	// bool
	dict.Store(true, true)
	// struct
	person := &Person{}
	person.name = "lala"
	person.age = 18
	dict.Store("person", person)
}

func gMap() {
	dict := make(map[string]string)
	// string
	dict["aaa"] = "AAA"
	// int
	// dict[1] = 1	// 会报错
}
