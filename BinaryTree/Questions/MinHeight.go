package main

import "fmt"

// 计算 当前树 的 最小深度：
func main() {
	// root1 最大 深度 为 3
	root1 := CreateTreeNode(5)
	root1.Left = CreateTreeNode(3)
	root1.Right = CreateTreeNode(8)
	root1.Left.Left = CreateTreeNode(1)
	root1.Left.Left.Left = CreateTreeNode(4)
	root1.Left.Left.Left.Left = CreateTreeNode(8)
	root1.Left.Left.Left.Left.Left = CreateTreeNode(3)
	root1.Left.Left.Left.Left.Left.Right = CreateTreeNode(7)
	root1.Right.Left = CreateTreeNode(7)
	root1.Right.Right = CreateTreeNode(9)
	fmt.Println(minHeight(root1))
}

func minHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := minHeight(root.Left), minHeight(root.Right)
	// 特判: 左右子树 有空树:
	if leftHeight == 0 || rightHeight == 0 {
		return leftHeight + rightHeight + 1
	} else {
		return min(leftHeight, rightHeight) + 1
	}
}
