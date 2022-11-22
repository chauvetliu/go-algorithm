package main

import "sort"

//
//合并两个有序数组
//题目：给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

// 合并两个数组后再排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 另开辟一个数组，双指针法排序两数组，最后再拷贝回 num1
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sort := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if i == m {
			sort = append(sort, nums2[j:]...)
			break
		}
		if j == n {
			sort = append(sort, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[j] {
			sort = append(sort, nums1[i])
			i++
		} else {
			sort = append(sort, nums2[j])
			j++
		}
	}
	copy(nums1, sort)
}

// 逆双指针法，倒序从两个数组中找到较大的一个放入 num1 中
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, tail := m-1, n-1, m+n-1
	for i >= 0 || j >= 0 {
		var cur int
		if i == -1 {
			cur = nums2[j]
			j--
		} else if j == -1 {
			cur = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			cur = nums2[j]
			j--
		} else {
			cur = nums1[i]
			i--
		}
		nums1[tail] = cur
		tail--
	}
}
