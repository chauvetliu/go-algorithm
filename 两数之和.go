package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}

func main() {
	var target = 9
	var nums = []int{15, 7, 11, 2, 1}
	fmt.Println(twoSum(nums, target))
}
