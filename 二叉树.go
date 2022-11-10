package main

import "fmt"

type TreeNode struct {
	Value     int       `json:"value"`
	LeftNode  *TreeNode `json:"left_node"`
	RightNode *TreeNode `json:"right_node"`
}

func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Value)
		traversal(node.LeftNode)
		traversal(node.RightNode)
	}
	traversal(root)
	return res
}

func main() {

	//前序遍历

	a := func() int {
		n3 := 2 + 9
		//fmt.Println(&n3)
		return n3
	} //传入参数

	fmt.Println(a)

	//定义一个匿名函数并调用
	res := func(n1 int, n2 int) int {
		n3 := n1 + n2
		return n3
	}(10, 10) //传入参数
	fmt.Println(res)

}
