package main

import "fmt"

type ListMerge struct {
	Num  int
	Next *ListMerge
}

// 递归方式1
func mergeTwoLists(list1 *ListMerge, list2 *ListMerge) *ListMerge {
	if list1 == nil || list2 == nil {
		return nil
	}

	if list1.Num < list2.Num {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}

// 方式二
func mergeTwoLists2(l1 *ListMerge, l2 *ListMerge) *ListMerge {
	res := &ListMerge{}
	temp := res
	for l1 != nil && l2 != nil {
		if l1.Num < l2.Num {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	} else {
		temp.Next = l2
	}
	return res.Next
}

// 方式二
func mergeTwoLists3(l1 *ListMerge, l2 *ListMerge) *ListMerge {
	res := &ListMerge{}
	temp := res
	for l1 != nil && l2 != nil {
		if l1.Num < l2.Num {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	return temp.Next
}

func main() {

	n1 := &ListMerge{
		Num: 1,
		Next: &ListMerge{
			Num: 3,
			Next: &ListMerge{
				Num: 5,
			},
		},
	}

	n2 := &ListMerge{
		Num: 2,
		Next: &ListMerge{
			Num: 4,
			Next: &ListMerge{
				Num: 6,
			},
		},
	}

	n := mergeTwoLists3(n1, n2)
	for n != nil {
		fmt.Printf("%d \t ", n.Num)
		n = n.Next
	}
	fmt.Println()
}
