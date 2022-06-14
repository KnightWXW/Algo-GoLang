package main

import (
	"fmt"
)

// 判断 当前树 是否为 平衡二叉树：
func main() {
	// root1 是 平衡二叉树
	root1 := CreateTreeNode(5)
	root1.Left = CreateTreeNode(3)
	root1.Right = CreateTreeNode(8)
	root1.Left.Left = CreateTreeNode(1)
	root1.Left.Right = CreateTreeNode(4)
	root1.Right.Left = CreateTreeNode(7)
	root1.Right.Right = CreateTreeNode(9)
	fmt.Println(isBT(root1))
	fmt.Println(is_BT(root1))

	// root2 不是 平衡二叉树
	root2 := CreateTreeNode(5)
	root2.Left = CreateTreeNode(3)
	root2.Right = CreateTreeNode(8)
	root2.Left.Left = CreateTreeNode(1)
	root2.Left.Right = CreateTreeNode(4)
	root2.Left.Left.Left = CreateTreeNode(9)
	fmt.Println(isBT(root2))
	fmt.Println(is_BT(root2))
}

//------------------------------普通递归------------------------------------------
func is_BT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(height(root.Left), height(root.Right)) + 1
	}
	return abs(height(root.Left)-height(root.Right)) <= 1
}

//-------------------------------------------------------------------------------

//-------------------------------递归套路------------------------------------------
func isBT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return processIsBT(root).Flag
}

func processIsBT(root *TreeNode) *InfoIsBT {
	if root == nil {
		return createInfoIsBT(true, 0)
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
	return createInfoIsBT(flag, height)
}

type InfoIsBT struct {
	Flag   bool
	Height int
}

func createInfoIsBT(flag bool, height int) *InfoIsBT {
	return &InfoIsBT{flag, height}
}

//----------------------------------------------------------------------------
