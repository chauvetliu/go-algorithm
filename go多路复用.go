package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		ch1 = make(chan string)
		ch2 = make(chan string)
	)
	go test2(ch1)
	go test3(ch2)

	select {
	case res := <-ch1:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	}
}

func test2(ch chan string) {
	fmt.Println("我是test2")
	time.Sleep(3 * time.Second) // 等待两个goroutine都执行完
	ch <- "test2的数据"
}
func test3(ch chan string) {
	fmt.Println("我是test3")
	time.Sleep(3 * time.Second) // 等待两个goroutine都执行完
	ch <- "test3的数据"
}
