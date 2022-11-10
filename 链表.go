package main

import "fmt"

type List struct {
	Node     string
	NextNode *List
}

func main() {

	var list1 List
	list1.Node = "liu"

	var list2 List
	list2.Node = "xiao"

	list1.NextNode = &list2

	var list3 List
	list3.Node = "wei"
	
	list2.NextNode = &list3

	each(&list1)
}

func each(node *List) {
	for node != nil {
		fmt.Println(*node)
		node = node.NextNode
	}
}
