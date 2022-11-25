package main

import "fmt"

type Tree struct {
	Name  string
	Left  *Tree
	Right *Tree
}

// 前序遍历：根结点 -> 左子树-> 右子树

func main() {

	var tree Tree
	tree.Name = "根树"

	//左树
	var leftTree Tree
	leftTree.Name = "左树"

	//根节点的左树
	tree.Left = &leftTree

	//右树
	var rightTree Tree
	rightTree.Name = "右树"

	//根节点的右树
	tree.Right = &rightTree

	Ergodic(&tree)
}

func Ergodic(eTree *Tree) {

	if eTree == nil {
		return
	}

	//输出根节点
	fmt.Println(eTree)

	//遍历左子树
	Ergodic(eTree.Left)

	//遍历右子树
	Ergodic(eTree.Right)
}
