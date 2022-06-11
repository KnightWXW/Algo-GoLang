package main

import (
	"fmt"
)

// 计算给定二叉树中最大的二叉搜索树（BST）子树，最大指的是子树的节点数最多
func main() {
	// 应为 6
	root1 := CreateTreeNode(4)
	root1.Left = CreateTreeNode(7)
	root1.Left.Left = CreateTreeNode(5)
	root1.Left.Right = CreateTreeNode(8)
	root1.Left.Right.Right = CreateTreeNode(9)
	root1.Left.Left.Left = CreateTreeNode(4)
	root1.Left.Left.Right = CreateTreeNode(6)
	root1.Right = CreateTreeNode(6)
	root1.Right.Left = CreateTreeNode(4)
	root1.Right.Right = CreateTreeNode(7)
	root1.Right.Right.Right = CreateTreeNode(9)
	fmt.Println(maxBSTSubtreeSize(root1))

	// 应为 4
	root2 := CreateTreeNode(4)
	root2.Left = CreateTreeNode(7)
	root2.Left.Left = CreateTreeNode(10)
	root2.Left.Right = CreateTreeNode(8)
	root2.Left.Right.Right = CreateTreeNode(9)
	root2.Left.Left.Left = CreateTreeNode(4)
	root2.Left.Left.Right = CreateTreeNode(6)
	root2.Right = CreateTreeNode(6)
	root2.Right.Left = CreateTreeNode(4)
	root2.Right.Right = CreateTreeNode(7)
	root2.Right.Right.Right = CreateTreeNode(9)
	fmt.Println(maxBSTSubtreeSize(root2))
}

func maxBSTSubtreeSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return processMaxBSTSubtreeSize(root).MaxBSTSize
}

func processMaxBSTSubtreeSize(root *TreeNode) *InfoMaxBSTSubtreeSize {
	if root == nil {
		return nil
	}
	leftInfo, rightInfo := processMaxBSTSubtreeSize(root.Left), processMaxBSTSubtreeSize(root.Right)
	// 更新 MaxVal, MinVal, Size:
	maxVal, minVal, size := root.Val, root.Val, 1
	if leftInfo != nil {
		maxVal = max(leftInfo.MaxVal, maxVal)
		minVal = max(leftInfo.MinVal, minVal)
		size += leftInfo.Size
	}
	if rightInfo != nil {
		maxVal = max(rightInfo.MaxVal, maxVal)
		minVal = max(rightInfo.MinVal, minVal)
		size += rightInfo.Size
	}
	// 更新 MaxBSTDistance:
	s1, s2, s3 := -1, -1, -1

	// 一、root 不是 最终 所要的节点：
	// 在左子树中寻找：
	if leftInfo != nil {
		s1 = leftInfo.MaxBSTSize
	}
	// 在右子树中寻找：
	if rightInfo != nil {
		s2 = rightInfo.MaxBSTSize
	}

	// 二、root 是 最终 所要的节点：
	// 判断 左右子树 是否 BST:
	leftIsBST, rightIsBST := false, false
	if leftInfo == nil || leftInfo.Size == leftInfo.MaxBSTSize {
		leftIsBST = true
	}
	if rightInfo == nil || rightInfo.Size == rightInfo.MaxBSTSize {
		rightIsBST = true
	}

	if leftIsBST && rightIsBST {
		// 判断 左子树的MaxVal < root.Val < 右子树的MinVal
		leftMaxLessVal, rightMinMoreVal := false, true
		if leftInfo == nil || leftInfo.MaxVal < root.Val {
			leftMaxLessVal = true
		}
		if rightInfo == nil || rightInfo.MaxVal > root.Val {
			rightMinMoreVal = true
		}
		// 更新 s3:
		if leftMaxLessVal && rightMinMoreVal {
			s3 = 1
			if leftInfo != nil {
				s3 += leftInfo.Size
			}
			if rightInfo != nil {
				s3 += rightInfo.Size
			}
		}
	}
	return createInfoMaxBSTSubtreeSize(max(max(s1, s2), s3), size, maxVal, minVal)
}

type InfoMaxBSTSubtreeSize struct {
	MaxBSTSize int
	Size       int
	MaxVal     int
	MinVal     int
}

func createInfoMaxBSTSubtreeSize(maxBSTSize, size, maxVal, minVal int) *InfoMaxBSTSubtreeSize {
	return &InfoMaxBSTSubtreeSize{maxBSTSize, size, maxVal, minVal}
}
