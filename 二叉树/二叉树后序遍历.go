package main

//后序遍历：左子树-> 右子树 -> 根结点

import "fmt"

// 定义结构体
type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student //左子树指针
	right *Student //右子树指针
}

// 二叉树定义
func main() {
	//根节点
	var root Student
	root.Name = "root"
	root.Age = 18
	root.Score = 88

	//一级左子树
	var left1 Student
	left1.Name = "left1"
	left1.Age = 20
	left1.Score = 80

	root.left = &left1

	//一级右子树
	var right1 Student
	right1.Name = "right1"
	right1.Age = 22
	right1.Score = 100

	root.right = &right1

	//调用遍历函数
	Req(&root)

}

// 递归算法遍历整个二叉树
func Req(tmp *Student) {
	if tmp == nil {
		return
	}

	//遍历左子树
	Req(tmp.left)

	//输出root节点
	fmt.Println(tmp)

	//遍历右子树
	Req(tmp.right)
}
