package main

//中序遍历：左子树-> 根结点 -> 右子树

import "fmt"

type Tree2 struct {
	Name  string
	Left  *Tree2
	Right *Tree2
}

// 前序遍历：根结点 -> 左子树-> 右子树

func main() {

	var tree Tree2
	tree.Name = "根树"

	//左树
	var leftTree Tree2
	leftTree.Name = "左树"

	//根节点的左树
	tree.Left = &leftTree

	//右树
	var rightTree Tree2
	rightTree.Name = "右树"

	//根节点的右树
	tree.Right = &rightTree

	Ergodic2(&tree)
}

func Ergodic2(eTree2 *Tree2) {

	if eTree2 == nil {
		return
	}

	//遍历左子树
	Ergodic2(eTree2.Left)

	//输出根节点
	fmt.Println(eTree2)

	//遍历右子树
	Ergodic2(eTree2.Right)
}
