package main

import (
	"fmt"
)

// 计算给定二叉树中最大的二叉搜索树（BST）子树，最大指的是子树的节点数最多, 返回该子树的头结点
func main() {

	root1 := CreateTreeNode(4)
	node1 := CreateTreeNode(7)
	node2 := CreateTreeNode(6)
	node3 := CreateTreeNode(5)
	node4 := CreateTreeNode(8)
	node5 := CreateTreeNode(4)
	node6 := CreateTreeNode(7)
	node7 := CreateTreeNode(4)
	node8 := CreateTreeNode(6)
	node9 := CreateTreeNode(9)
	node10 := CreateTreeNode(9)

	root1.Left = node1
	root1.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	node3.Left = node7
	node3.Right = node8
	node4.Right = node9
	node6.Right = node10

	fmt.Println(maxBSTSubtreeHead(root1)) //  node1---(val: 7)

	root2 := CreateTreeNode(4)
	node1_ := CreateTreeNode(10)
	node2_ := CreateTreeNode(6)
	node3_ := CreateTreeNode(5)
	node4_ := CreateTreeNode(8)
	node5_ := CreateTreeNode(4)
	node6_ := CreateTreeNode(7)
	node7_ := CreateTreeNode(4)
	node8_ := CreateTreeNode(6)
	node9_ := CreateTreeNode(9)
	node10_ := CreateTreeNode(9)

	root2.Left = node1_
	root2.Right = node2_
	node1_.Left = node3_
	node1_.Right = node4_
	node2_.Left = node5_
	node2_.Right = node6_
	node3_.Left = node7_
	node3_.Right = node8_
	node4_.Right = node9_
	node6_.Right = node10_

	fmt.Println(maxBSTSubtreeHead(root2)) //  node2---(val: 6)

}

func maxBSTSubtreeHead(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return processMaxBSTSubtreeHead(root).MaxBSTHead
}

func processMaxBSTSubtreeHead(root *TreeNode) *InfoMaxBSTSubtreeHead {
	if root == nil {
		return nil
	}
	leftInfo, rightInfo := processMaxBSTSubtreeHead(root.Left), processMaxBSTSubtreeHead(root.Right)

	// 更新 MaxVal, MinVal, maxBSTHead, maxBSTSize:
	var maxBSTHead *TreeNode
	maxBSTSize, maxVal, minVal := 0, root.Val, root.Val
	if leftInfo != nil {
		maxVal = max(leftInfo.MaxVal, maxVal)
		minVal = max(leftInfo.MinVal, minVal)
		maxBSTSize = leftInfo.MaxBSTSize
		maxBSTHead = leftInfo.MaxBSTHead
	}
	if rightInfo != nil {
		maxVal = max(rightInfo.MaxVal, maxVal)
		minVal = max(rightInfo.MinVal, minVal)
		if rightInfo.MaxBSTSize > maxBSTSize {
			maxBSTSize = rightInfo.MaxBSTSize
			maxBSTHead = rightInfo.MaxBSTHead
		}
	}
	// 判断 左右子树 是否 BST:
	leftIsBST, rightIsBST := false, false
	if leftInfo == nil || leftInfo.MaxBSTHead == root.Left {
		leftIsBST = true
	}
	if rightInfo == nil || rightInfo.MaxBSTHead == root.Right {
		rightIsBST = true
	}
	// 判断 左子树的MaxVal < root.Val < 右子树的MinVal
	leftMaxLessVal, rightMinMoreVal := false, false
	if leftInfo == nil || leftInfo.MaxVal < root.Val {
		leftMaxLessVal = true
	}
	if rightInfo == nil || rightInfo.MinVal > root.Val {
		rightMinMoreVal = true
	}

	if leftIsBST && rightIsBST && leftMaxLessVal && rightMinMoreVal {
		maxBSTHead = root
		if leftInfo != nil {
			maxBSTSize += leftInfo.MaxBSTSize
		}
		if rightInfo != nil {
			maxBSTSize += rightInfo.MaxBSTSize
		}
		maxBSTSize++
	}
	return createInfoMaxBSTSubtreeHead(maxBSTHead, maxBSTSize, maxVal, minVal)
}

/*
if ((leftInfo == null ? true : (leftInfo.maxSubBSTHead == X.left && leftInfo.max < X.value))
&& (rightInfo == null ? true : (rightInfo.maxSubBSTHead == X.right && rightInfo.min > X.value))) {
maxSubBSTHead = X;
maxSubBSTSize = (leftInfo == null ? 0 : leftInfo.maxSubBSTSize)
+ (rightInfo == null ? 0 : rightInfo.maxSubBSTSize) + 1;
}
*/

type InfoMaxBSTSubtreeHead struct {
	MaxBSTHead *TreeNode
	MaxBSTSize int
	MaxVal     int
	MinVal     int
}

func createInfoMaxBSTSubtreeHead(maxBSTHead *TreeNode, maxBSTSize, maxVal, minVal int) *InfoMaxBSTSubtreeHead {
	return &InfoMaxBSTSubtreeHead{maxBSTHead, maxBSTSize, maxVal, minVal}
}
