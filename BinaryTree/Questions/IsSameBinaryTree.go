package main

import "fmt"

// 判断 当前 两棵树 是否 相同：
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

	root2 := CreateTreeNode(5)
	root2.Left = CreateTreeNode(3)
	root2.Right = CreateTreeNode(3)
	root2.Left.Left = CreateTreeNode(1)
	root2.Left.Right = CreateTreeNode(2)
	root2.Right.Left = CreateTreeNode(2)
	root2.Right.Right = CreateTreeNode(1)
	root2.Left.Left.Left = CreateTreeNode(8)
	root2.Right.Right.Right = CreateTreeNode(8)
	root2.Left.Right.Right = CreateTreeNode(6)
	root2.Right.Left.Left = CreateTreeNode(6)
	fmt.Println(isSameBinaryTree(root1, root2))
}

func isSameBinaryTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	return dfsIsSameBinaryTree(root1, root2)
}

func dfsIsSameBinaryTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val && dfsIsSameBinaryTree(root1.Left, root2.Left) == true && dfsIsSameBinaryTree(root1.Right, root2.Right) == true {
		return true
	} else {
		return false
	}
}
