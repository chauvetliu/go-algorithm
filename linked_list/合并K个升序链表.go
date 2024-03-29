package main

type ListNodeMerge struct {
	Val  int
	Next *ListNodeMerge
}

func mergeKLists(lists []*ListNodeMerge) *ListNodeMerge {
	dummy := &ListNodeMerge{}
	for _, v := range lists {
		dummy.Next = merge(dummy.Next, v)
	}
	return dummy.Next
}

func merge(l1, l2 *ListNodeMerge) *ListNodeMerge {
	dummy := &ListNodeMerge{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	return dummy.Next
}
