package main

import (
	"fmt"
)

// 判断 当前树 是否为 完全二叉树：
func main() {
	// root1 是 完全二叉树
	root1 := CreateTreeNode(5)
	root1.Left = CreateTreeNode(3)
	root1.Right = CreateTreeNode(8)
	root1.Left.Left = CreateTreeNode(1)
	root1.Left.Right = CreateTreeNode(4)
	root1.Right.Left = CreateTreeNode(7)
	root1.Right.Right = CreateTreeNode(9)
	fmt.Println(isCBT(root1))

	// root2 不是 完全二叉树
	root2 := CreateTreeNode(5)
	root2.Left = CreateTreeNode(3)
	root2.Right = CreateTreeNode(8)
	root2.Left.Left = CreateTreeNode(10)
	root2.Left.Right = CreateTreeNode(4)
	root2.Right.Right = CreateTreeNode(9)
	fmt.Println(isCBT(root2))
}

func isCBT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return processIsCBT(root).IsCBT
}

func processIsCBT(root *TreeNode) *InfoIsCBT {
	if root == nil {
		return createInfoIsCBT(true, true, 0)
	}
	leftInfo, rightInfo := processIsCBT(root.Left), processIsCBT(root.Right)

	// 更新 Height：
	height := max(leftInfo.Height, rightInfo.Height) + 1

	// 更新 IsFBT：
	isFBT := true
	if leftInfo.IsFBT == false || rightInfo.IsFBT == false || leftInfo.Height != rightInfo.Height {
		isFBT = false
	}

	// 更新 IsCBT：
	isCBT := false
	// 若 root 为 满二叉树，必为 完全二叉树
	if isFBT == true {
		isCBT = true
	} else {
		// 若 root 左右子树为 完全二叉树，则 左子树树高 == 右子树树高 + 1
		if leftInfo.IsCBT == true && rightInfo.IsCBT == true && leftInfo.Height == rightInfo.Height+1 {
			isCBT = true
			// 若 root 左右子树为 满二叉树，则 左子树树高 == 右子树树高 + 1
		} else if leftInfo.IsFBT == true && rightInfo.IsFBT == true && leftInfo.Height == rightInfo.Height+1 {
			isCBT = true
			// 若 root 左子树为 满二叉树，右子树为 完全二叉树， 则 左子树树高 == 右子树树高 + 1
		} else if leftInfo.IsFBT == true && rightInfo.IsCBT == true && leftInfo.Height == rightInfo.Height+1 {
			isCBT = true
		}
	}

	return createInfoIsCBT(isCBT, isFBT, height)
}

type InfoIsCBT struct {
	IsCBT  bool
	IsFBT  bool
	Height int
}

func createInfoIsCBT(isCBT, isFBT bool, height int) *InfoIsCBT {
	return &InfoIsCBT{isCBT, isFBT, height}
}
