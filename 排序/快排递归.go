package main

import (
	"fmt"
)

// QuickSort 递归
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]          //第一个数据为基准
	low := make([]int, 0, 0)     //比我小的数据
	hight := make([]int, 0, 0)   //比我大的数据
	mid := make([]int, 0, 0)     //与我一样大的数据
	mid = append(mid, splitdata) //加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight) //切割递归处理
	myarr := append(append(low, mid...), hight...)
	return myarr
}

// 快读排序算法
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}

/*
*
快速排序
//*/
//func quickSort(data []int) {
//	if len(data) < 1 {
//		return
//	}
//	l, r := 0, len(data)-1
//	//基准值
//	base := data[0]
//	for i := 0; i <= r; {
//		//比基准值大的放右边
//		if data[i] > base {
//			data[r], data[i] = data[i], data[r]
//			r--
//		} else {
//			//比基准值小或等于的放左边
//			data[l], data[i] = data[i], data[l]
//			l++
//			i++
//		}
//	}
//	quickSort(data[:l])
//	quickSort(data[l+1:])
//}
