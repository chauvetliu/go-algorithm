package main

import "fmt"

//func twoSum(nums []int, target int) []int {
//	for k1, _ := range nums {
//		for k2 := k1 + 1; k2 < len(nums); k2++ {
//			if target == nums[k1] + nums[k2] {
//				return []int{k1, k2}
//			}
//		}
//	}
//	return []int{}
//}

//func twoSum(nums []int, target int) []int {
//	for k1, _ := range nums {
//		for k2 := k1 + 1; k2 < len(nums); k2++ {
//			if target == nums[k1]+nums[k2] {
//				return []int{k1, k2}
//			}
//		}
//	}
//	return []int{}
//}

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

	//var target = 9
	//var val = 8
	//
	//m := make(map[int]int)
	//preIndex, ok := m[target-val]
	//fmt.Println(preIndex)
	//fmt.Println(ok)

	fmt.Println(twoSum([]int{15, 7, 11, 2, 1}, 9))
}
