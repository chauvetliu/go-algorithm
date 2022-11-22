package main

import "fmt"

// 数组升序排序
func sortArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 6, 2, 3, 7, 4, 8, 5}
	sArr := sortArray(arr)
	fmt.Println(sArr)
}
