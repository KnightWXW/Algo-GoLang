package main

import (
	"fmt"
)

// 计算 当前树 的 最大距离：
func main() {
	root := CreateTreeNode(5)
	root.Left = CreateTreeNode(3)
	root.Left.Left = CreateTreeNode(1)
	root.Left.Left.Left = CreateTreeNode(4)
	root.Left.Left.Left.Left = CreateTreeNode(-1)
	root.Right = CreateTreeNode(8)
	root.Right.Right = CreateTreeNode(6)
	root.Right.Right.Right = CreateTreeNode(9)
	root.Right.Right.Right.Right = CreateTreeNode(7)
	fmt.Println(maxDistance(root))
}

func maxDistance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return processMaxDistance(root).MaxDistance
}

func processMaxDistance(root *TreeNode) *InfoMaxDistance {
	if root == nil {
		return createInfoMaxDistance(0, 0)
	}
	leftInfo, rightInfo := processMaxDistance(root.Left), processMaxDistance(root.Right)
	// 更新 Height：
	height := max(leftInfo.Height, rightInfo.Height) + 1
	// 更新 MaxDistance：
	maxDistance := max(leftInfo.Height+rightInfo.Height+1, max(leftInfo.MaxDistance, rightInfo.MaxDistance))

	return createInfoMaxDistance(maxDistance, height)
}

type InfoMaxDistance struct {
	MaxDistance int
	Height      int
}

func createInfoMaxDistance(maxDistance, height int) *InfoMaxDistance {
	return &InfoMaxDistance{maxDistance, height}
}
