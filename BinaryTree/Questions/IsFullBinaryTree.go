package main

import (
	"fmt"
	"math"
)

// 判断 当前树 是否为 满二叉树：
func main() {
	// root1 是 满二叉树
	root1 := CreateTreeNode(5)
	root1.Left = CreateTreeNode(3)
	root1.Right = CreateTreeNode(8)
	root1.Left.Left = CreateTreeNode(1)
	root1.Left.Right = CreateTreeNode(4)
	root1.Right.Left = CreateTreeNode(7)
	root1.Right.Right = CreateTreeNode(9)
	fmt.Println(isFBT1(root1))
	fmt.Println(isFBT2(root1))

	// root2 不是 满二叉树
	root2 := CreateTreeNode(5)
	root2.Left = CreateTreeNode(3)
	root2.Right = CreateTreeNode(8)
	root2.Left.Left = CreateTreeNode(1)
	root2.Left.Right = CreateTreeNode(4)
	fmt.Println(isFBT1(root2))
	fmt.Println(isFBT2(root2))
}

//------------------------------方法一---------------------------------------

func isFBT1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	info := processIsFBT1(root)
	return info.Number == int(math.Pow(float64(2), float64(info.Height)))-1
}

func processIsFBT1(root *TreeNode) *InfoIsFBT1 {
	if root == nil {
		return createInfoIsFBT1(0, 0)
	}
	leftInfo, rightInfo := processIsFBT1(root.Left), processIsFBT1(root.Right)

	// 更新 Height：
	height := max(leftInfo.Height, rightInfo.Height) + 1
	// 更新 Number：
	number := leftInfo.Number + rightInfo.Number + 1

	return createInfoIsFBT1(height, number)
}

type InfoIsFBT1 struct {
	Height int
	Number int
}

func createInfoIsFBT1(height, number int) *InfoIsFBT1 {
	return &InfoIsFBT1{height, number}
}

//------------------------------方法二---------------------------------------

func isFBT2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return processIsFBT2(root).Flag
}

func processIsFBT2(root *TreeNode) *InfoIsFBT2 {
	if root == nil {
		return createInfoIsFBT2(true, 0)
	}
	leftInfo, rightInfo := processIsFBT2(root.Left), processIsFBT2(root.Right)

	// 更新 Flag：
	flag := true
	if leftInfo.Flag == false || rightInfo.Flag == false {
		flag = false
	}
	if leftInfo.Height != rightInfo.Height {
		flag = false
	}

	// 更新 Height：
	height := max(leftInfo.Height, rightInfo.Height) + 1

	return createInfoIsFBT2(flag, height)
}

type InfoIsFBT2 struct {
	Flag   bool
	Height int
}

func createInfoIsFBT2(flag bool, height int) *InfoIsFBT2 {
	return &InfoIsFBT2{flag, height}
}
