package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 快速排序
// 1、left表示数组左边的下标
// 2、right表示数组右边的下标
func QuickSort2(left, right int, array *[80000]int) {
	l := left
	r := right

	pivot := array[(left+right)/2]
	//for循环的目标是将比pivot小的数放在左边，比pivot大的数放在右边
	for l < r {
		//从pivot左边找到一个大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从pivot右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		//表明本次分解任务完成
		if l >= r {
			break
		}
		//交换
		array[l], array[r] = array[r], array[l]
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
		//	fmt.Printf("%v\n", *array)
	}
	//可注释以下代码做测试
	//如果l==r，在移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort2(left, r, array)
	}
	//向右递归
	if right > l {
		QuickSort2(l, right, array)
	}
}

// 快读排序算法
func main() {
	//arr := [11]int{-1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//slcie := arr[:]
	//fmt.Printf("%T\n", slcie)

	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}
	start := time.Now().Unix()
	//fmt.Println("初始", arr)
	QuickSort2(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	fmt.Printf("快排耗时：%ds", end-start)
	//fmt.Println("main", arr)

}
