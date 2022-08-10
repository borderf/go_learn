package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

var fileChan = make(chan string, 10000)
var wg sync.WaitGroup

// var readFinish = make(chan struct{})
var writeFinish = make(chan struct{})

func readFile(fileName string) {
	defer wg.Done()
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	// 构建FileReader
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 { // 文件最后一行不是空行
					fileChan <- line
				}
				break
			} else {
				fmt.Println(err.Error())
				break
			}
		} else {
			fileChan <- line
		}

	}
}

func writeFile(fileName string) {
	defer close(writeFinish)
	fout, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)
	// LOOP:
	// 	for {
	// 		select {
	// 		case <-readFinish:
	// 			close(fileChan)
	// 			for line := range fileChan {
	// 				writer.WriteString(line)
	// 			}
	// 			break LOOP
	// 		case line := <-fileChan:
	// 			writer.WriteString(line)
	// 		}
	// 	}
	for {
		if line, ok := <-fileChan; ok { // ok为true，表面管道里面还有内容，false则代表从管道里拿不到内容
			writer.WriteString(line)
		} else {
			// fileChan 已关闭
			break
		}
	}
	writer.Flush()
}

func main() {
	wg.Add(3)
	for i := 0; i < 3; i++ {
		fileName := "goroutine_learn/homework/file/" + strconv.Itoa(i+1)
		go readFile(fileName)
	}
	go writeFile("goroutine_learn/homework/file/merge")
	wg.Wait()
	// close(readFinish)
	close(fileChan)
	<-writeFinish
}
