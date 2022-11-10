package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode = nil
//	var curr = head
//	if curr != nil {
//		var next = curr.Next
//		curr.Next = prev
//		prev = curr
//		curr = next
//	}
//	return prev
//}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr = head
	if curr != nil {
		//获取下一个节点
		var next = curr.Next

		//下个节点等于上个节点
		curr.Next = prev

		//当前节点复制给前一节点
		prev = curr

		//将temp节点赋值给当前节点，用于继续往下执行
		curr = next

	}

	return prev
}

func main() {

	var list ListNode
	fmt.Println(reverseList(list.Next))

}
