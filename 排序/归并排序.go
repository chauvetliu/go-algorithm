package main

//归并排序将数组分成两个子数组分别排序，并将有序的子数组归并以将整个数组排序

/*
*
归并排序
*/
func mergeSort(data []int) (result []int) {

	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	//分治排序左子数组
	left := mergeSort(data[:mid])
	//分治排序右子数组
	right := mergeSort(data[mid:])
	//分开排序后合并成一个新数组，此时原数组已经有序
	return merge(left, right)
}

func merge(left []int, right []int) (result []int) {

	l, r := 0, 0
	//排序，左右两边比较，放入result
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
