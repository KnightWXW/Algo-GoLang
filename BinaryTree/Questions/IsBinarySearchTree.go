package main

import (
	"fmt"
)

// 判断 当前树 是否为 二叉搜索树：
func main() {
	// root1 是 二叉搜索树
	root1 := CreateTreeNode(5)
	root1.Left = CreateTreeNode(3)
	root1.Right = CreateTreeNode(8)
	root1.Left.Left = CreateTreeNode(1)
	root1.Left.Right = CreateTreeNode(4)
	root1.Right.Left = CreateTreeNode(7)
	root1.Right.Right = CreateTreeNode(9)
	fmt.Println(isBST(root1))

	// root2 不是 二叉搜索树
	root2 := CreateTreeNode(5)
	root2.Left = CreateTreeNode(3)
	root2.Right = CreateTreeNode(8)
	root2.Left.Left = CreateTreeNode(10)
	root2.Left.Right = CreateTreeNode(4)
	root2.Right.Left = CreateTreeNode(7)
	root2.Right.Right = CreateTreeNode(9)
	fmt.Println(isBST(root2))
}

func isBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return processIsBST(root).Flag
}

func processIsBST(root *TreeNode) *InfoIsBST {
	if root == nil {
		return nil
	}
	leftInfo, rightInfo := processIsBST(root.Left), processIsBST(root.Right)
	flag, maxVal, minVal := true, root.Val, root.Val

	// 更新 MaxVal：
	if leftInfo != nil {
		maxVal = max(leftInfo.MaxVal, maxVal)
	}
	if rightInfo != nil {
		maxVal = max(rightInfo.MaxVal, maxVal)
	}

	// 更新 MinVal：
	if leftInfo != nil {
		minVal = min(leftInfo.MaxVal, minVal)
	}
	if rightInfo != nil {
		minVal = min(rightInfo.MinVal, minVal)
	}

	// 更新 Flag：
	if leftInfo != nil && leftInfo.Flag == false {
		flag = false
	}
	if rightInfo != nil && rightInfo.Flag == false {
		flag = false
	}
	if leftInfo != nil && leftInfo.MaxVal >= root.Val {
		flag = false
	}
	if rightInfo != nil && rightInfo.MinVal <= root.Val {
		flag = false
	}
	return createInfoIsBST(flag, maxVal, minVal)
}

type InfoIsBST struct {
	Flag   bool
	MaxVal int
	MinVal int
}

func createInfoIsBST(flag bool, maxVal, minVal int) *InfoIsBST {
	return &InfoIsBST{flag, maxVal, minVal}
}
