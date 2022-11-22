package main

type ListNodeII struct {
	Val  int
	Next *ListNodeII
}

func reverseBetween(head *ListNodeII, left int, right int) *ListNodeII {
	// 定义一个哨兵结点，方便后续处理
	headPre := &ListNodeII{Val: 0, Next: head}
	// 先找到第 m 个结点的前一个结点，即第一部分的尾部结点
	firstTail := headPre
	for i := 1; i < left; i++ {
		firstTail = firstTail.Next
	}
	// 记录第二部分翻转后的尾部结点
	secondTail := firstTail.Next
	// 将接下来 right - left + 1 个结点用头插法插入到 firstTail 后面
	cur := firstTail.Next
	for i := right - left; i >= 0; i-- {
		// 先保存下一个结点
		next := cur.Next
		// 将当前结点插入到 firstTail 后面
		cur.Next = firstTail.Next
		firstTail.Next = cur

		cur = next
	}
	// 将剩余部分挂在第二部分翻转后的尾部结点 secondTail 后面
	secondTail.Next = cur

	return headPre.Next
}

func main() {

}