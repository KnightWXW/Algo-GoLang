package main

import "fmt"

// 判断 当前树 是否是 镜像二叉树：
func main() {
	root1 := CreateTreeNode(5)
	root1.Left = CreateTreeNode(3)
	root1.Right = CreateTreeNode(3)
	root1.Left.Left = CreateTreeNode(1)
	root1.Left.Right = CreateTreeNode(2)
	root1.Right.Left = CreateTreeNode(2)
	root1.Right.Right = CreateTreeNode(1)
	root1.Left.Left.Left = CreateTreeNode(8)
	root1.Right.Right.Right = CreateTreeNode(8)
	root1.Left.Right.Right = CreateTreeNode(6)
	root1.Right.Left.Left = CreateTreeNode(6)
	fmt.Println(isMirrorBinaryTree(root1))
}

func isMirrorBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsIsMirrorBinaryTree(root.Left, root.Right)
}

func dfsIsMirrorBinaryTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val && dfsIsMirrorBinaryTree(root1.Left, root2.Right) == true && dfsIsMirrorBinaryTree(root1.Right, root2.Left) == true {
		return true
	} else {
		return false
	}
}
