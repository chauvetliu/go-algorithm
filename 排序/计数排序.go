package main

//计数排序
//稳定排序
//其空间复杂度和时间复杂度均为O(n+k)线性时间复杂度，其中k是整数的范围（取决于辅助数组大小）
//非比较排序
//计数排序其实是桶数取 max - min + 1最大时的桶排序

func countSort(data []int) {

	min := getMin(data)
	max := getMax(data)
	//建立辅助数组
	c := make([]int, max-min+1)
	for i := range data {
		//以跟最小值的差为下标计数
		j := data[i] - min
		c[j]++
	}
	k := 0
	//恢复数组
	for i := range c {
		for c[i] != 0 {
			data[k] = i + min
			k++
			c[i]--
		}
	}
}

func getMax(data []int) (m int) {
	max := 0
	for i := 0; i < len(data); i++ {
		if max < data[i] {
			max = data[i]
		}
	}
	return max
}

func getMin(data []int) (m int) {
	min := 0
	for i := 0; i < len(data); i++ {
		if min > data[i] {
			min = data[i]
		}
	}
	return min
}
