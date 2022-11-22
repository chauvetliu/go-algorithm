package main

// 演示使用goroutine 顺序打印数字: 0,1,2,3,4,5,6,7,8,9

import (
	"fmt"
	"sync"
)

func main() {
	n := 10

	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)

	go func() {
		display2(n, ch)
		wg.Done()
	}()

	go func() {
		display2(n, ch)
		wg.Done()
	}()

	ch <- 0
	wg.Wait()
}

// display2 使用无缓冲通道开启两个 goroutine 之间来回交换打印叠加数字的思路
func display2(n int, ch chan int) {
	for {
		c, ok := <-ch
		if !ok {
			return
		}

		if c == n {
			close(ch)
			return
		}
		fmt.Printf("ch: %d \n", c)
		c++
		ch <- c
	}
}
