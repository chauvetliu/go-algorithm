package main

// 在乱序的条件下，时间复杂度是O(N^2)，有序条件下时间复杂度是O(N)
// 稳定排序
// 空间复杂度：O(1)
func bubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				//交换
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
