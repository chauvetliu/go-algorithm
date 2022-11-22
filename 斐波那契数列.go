package main

import "fmt"

// 用go实现斐波那契数列 又称黄金分割数列
func fib(num int) int {
	var result int
	if num < 2 {
		result = num
	} else {
		result = fib(num-1) + fib(num-2)
	}
	return result
}

// 费波那契数列由0和1开始，之后的费波那契系数就是由之前的两数相加而得出。
func main() {
	nu := 7
	var arr []int
	for i := 0; i <= nu; i++ {
		arr = append(arr, fib(i))
	}
	for _, j := range arr {
		fmt.Print(j)
		fmt.Print("  ")
	}
}
