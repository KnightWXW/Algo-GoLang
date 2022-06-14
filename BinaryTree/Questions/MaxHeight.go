package main

import "fmt"

// 计算 当前树 的 最大深度：
func main() {
	// root1 最大 深度 为 7
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
	fmt.Println(maxHeight(root1))
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := maxHeight(root.Left), maxHeight(root.Right)
	return max(leftHeight, rightHeight) + 1
}
