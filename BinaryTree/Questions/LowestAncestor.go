package main

import (
	"fmt"
)

// 计算 两个节点 的 最近公共祖先：
func main() {
	root := CreateTreeNode(0)
	node1 := CreateTreeNode(1)
	node2 := CreateTreeNode(2)
	node3 := CreateTreeNode(3)
	node4 := CreateTreeNode(4)
	node5 := CreateTreeNode(5)
	node6 := CreateTreeNode(6)
	node7 := CreateTreeNode(7)
	node8 := CreateTreeNode(8)
	node9 := CreateTreeNode(9)

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	node3.Left = node7
	node3.Right = node8
	node4.Left = node9

	fmt.Println(lowestAncestor(root, node5, node6))  //  node2
	fmt.Println(lowestAncestor_(root, node5, node6)) //  node2
	fmt.Println(lowestAncestor(root, node7, node9))  //  node1
	fmt.Println(lowestAncestor_(root, node7, node9)) //  node2
}

// 普通递归：
func lowestAncestor_(root, a, b *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == a.Val || root.Val == b.Val {
		return root
	}
	nodeA := lowestAncestor_(root.Left, a, b)
	nodeB := lowestAncestor_(root.Right, a, b)

	if nodeA == nil {
		return nodeB
	}
	if nodeB == nil {
		return nodeA
	}
	return root
}

func lowestAncestor(root, a, b *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return processLowestAncestor(root, a, b).Node
}

func processLowestAncestor(root, a, b *TreeNode) *InfoLowestAncestor {
	if root == nil {
		return createInfoLowestAncestor(false, false, nil)
	}
	leftInfo, rightInfo := processLowestAncestor(root.Left, a, b), processLowestAncestor(root.Right, a, b)

	// 更新 flagA：
	flagA := false
	if leftInfo.FlagA == true || rightInfo.FlagA == true || root == a {
		flagA = true
	}

	// 更新 flagB：
	flagB := false
	if leftInfo.FlagB == true || rightInfo.FlagB == true || root == b {
		flagB = true
	}

	// 更新 Node：
	var node *TreeNode
	if leftInfo.Node != nil {
		node = leftInfo.Node
	} else if rightInfo.Node != nil {
		node = rightInfo.Node
	} else {
		if flagA && flagB {
			node = root
		}
	}

	return createInfoLowestAncestor(flagA, flagB, node)
}

type InfoLowestAncestor struct {
	FlagA bool
	FlagB bool
	Node  *TreeNode
}

func createInfoLowestAncestor(flagA, flagB bool, node *TreeNode) *InfoLowestAncestor {
	return &InfoLowestAncestor{flagA, flagB, node}
}
