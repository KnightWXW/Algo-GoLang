package main

import (
	"strconv"
)

// 二叉树节点 TreeNode：
type TreeNode struct {
	Val   int       `二叉树节点的数值`
	Left  *TreeNode `二叉树节点的左子树节点`
	Right *TreeNode `二叉树节点的右子树节点`
}

// 二叉树 Tree：
type BinaryTree struct {
	root *TreeNode `二叉树的根节点`
	cnt  int       `二叉树的节点数量`
}

// 创建二叉树节点
func CreateTreeNode(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

// 向当前二叉树节点中添加二叉树节点,按照 左小右大 的原则
func (node *TreeNode) AddTreeNodeRandom(addNode *TreeNode) {
	if node.Val > addNode.Val {
		if node.Left == nil {
			node.Left = addNode
		} else {
			node.Left.AddTreeNodeRandom(addNode)
		}
	} else {
		if node.Right == nil {
			node.Right = addNode
		} else {
			node.Right.AddTreeNodeRandom(addNode)
		}
	}
}

// 向二叉树中添加二叉树节点
func (t *BinaryTree) Add(v int) {
	node := CreateTreeNode(v)
	if t.root == nil {
		t.root = node
	} else {
		t.root.AddTreeNodeRandom(node)
	}
	t.cnt++
}

// 返回二叉树的节点数量
func (t *BinaryTree) Size() int {
	return t.cnt
}

// 判断二叉树是否为空数：
func (t *BinaryTree) IsEmpty() bool {
	return t.cnt == 0
}

// 根据数组的顺序创建二叉树：
func AddTreeNodeAsArray(arr []string, i int) *TreeNode {
	var ans *TreeNode
	if i >= len(arr) {
		return nil
	}
	if arr[i] == "#" {
		return ans
	} else {
		v, _ := strconv.Atoi(arr[i])
		ans = CreateTreeNode(v)
		ans.Left = AddTreeNodeAsArray(arr, 2*i+1)
		ans.Right = AddTreeNodeAsArray(arr, 2*i+2)
	}
	return ans
}
