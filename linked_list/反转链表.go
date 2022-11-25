package main

import "fmt"

type ListNode struct {
	Num  int
	Next *ListNode
}

func main() {
	var listNode = &ListNode{
		Num: 1,
		Next: &ListNode{
			Num: 2,
			Next: &ListNode{
				Num: 3,
				Next: &ListNode{
					Num: 4,
					Next: &ListNode{
						Num:  5,
						Next: nil,
					},
				},
			},
		},
	}
	a := reverse(listNode)
	for a != nil {
		fmt.Println(a.Num)
		a = a.Next
	}
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//func reverse(head *ListNode) *ListNode {
//	var prev *ListNode = nil
//	var curr = head
//	if curr != nil {
//		//获取下一个节点
//		var next = curr.Next
//		//下个节点等于上个节点
//		curr.Next = prev
//		//当前节点复制给前一节点
//		prev = curr
//		//将temp节点赋值给当前节点，用于继续往下执行
//		curr = next
//	}
//
//	return prev
//}
