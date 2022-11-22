package main

//首先找到数组中的最小元素，然后将这个最小元素和数组的第一个元素交换位置，如果第一个元素就是最小元素，就和自己交换位置；
//再次，在剩下的元素中找到最小元素和数组中的第二个元素交换位置，如此往复，直到将整个数组排序，一句话总结就是，不断在剩余元素中找最小元素

//不稳定排序
//时间复杂度：O(n^2)
//空间复杂度：O(1)
//选择排序是数据移动最少的，每次交换两会改变两个数组元素的值，因此选择排序用了N次交换，交换次数和数组大小是线性关系

// 选择排序
func selectSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		//每次循环找出 i + 1到数组最后一个元素的这个区间的最小值，然后最小值和当前元素data[i]比较
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
}
