package main

// 堆排序是基于二叉堆（优先队列）实现的，通过建立最大堆或者最小堆，
// 删除堆顶的最值元素后将堆顶元素放入指定集合，重新建堆，最后集合便是有序序列
// 堆排序
func heapSort(data []int) {
	m := len(data)
	s := m / 2
	for i := s; i >= 0; i-- {
		heap(data, i, m-1)
	}
	for i := m - 1; i > 0; i-- {
		data[i], data[0] = data[0], data[i]
		heap(data, 0, i-1)
	}
}

// 以i为根节点建堆
func heap(data []int, i int, end int) {
	//左子节点
	l := 2*i + 1
	if l > end {
		return
	}
	n := l
	//右子节点
	r := 2*i + 2
	//当前最大元素是右子节点
	if r <= end && data[r] > data[n] {
		n = r
	}
	if data[i] > data[n] {
		return
	}
	//比根节点大则交换
	data[i], data[n] = data[n], data[i]
	//往下建堆
	heap(data, n, end)
}
