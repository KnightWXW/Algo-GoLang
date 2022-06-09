package main

import (
	"fmt"
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
	fmt.Println(isBT(root1))

	// root2 不是 满二叉树
	root2 := CreateTreeNode(5)
	root2.Left = CreateTreeNode(3)
	root2.Right = CreateTreeNode(8)
	root2.Left.Left = CreateTreeNode(1)
	root2.Left.Right = CreateTreeNode(4)
	root2.Left.Left.Left = CreateTreeNode(9)
	fmt.Println(isBT(root2))
}

func isFT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return processIsFT(root).Flag
}

func processIsFT(root *TreeNode) *InfoIsFT {
	if root == nil {
		return createInfoIsFT(true, 0)
	}
	leftInfo, rightInfo := processIsBT(root.Left), processIsBT(root.Right)
	flag, height := true, max(leftInfo.Height, rightInfo.Height)+1

	// 更新 Flag：
	if leftInfo.Flag == false {
		flag = false
	}
	if rightInfo.Flag == false {
		flag = false
	}
	if abs(leftInfo.Height-rightInfo.Height) >= 2 {
		flag = false
	}

	return createInfoIsFT(flag, height)
}

type InfoIsFT struct {
	Flag   bool
	Height int
}

func createInfoIsFT(flag bool, height int) *InfoIsFT {
	return &InfoIsFT{flag, height}
}
