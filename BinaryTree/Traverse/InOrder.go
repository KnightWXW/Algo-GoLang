package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []string{"5", "2", "7", "1", "0", "6", "8", "4", "3", "9"}
	var tree BinaryTree
	for i := 0; i < len(arr); i++ {
		v, _ := strconv.Atoi(arr[i])
		tree.Add(v)
	}
	fmt.Println("二叉树中序遍历(递归形式):")
	inOrderRecursion(tree.root)
	fmt.Println()
	fmt.Println("二叉树中序遍历(非递归形式):")
	inOrderIteration(tree.root)
}

// 中序遍历 的 递归实现:
func inOrderRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	inOrderRecursion(node.Left)
	fmt.Printf("%v ", node.Val)
	inOrderRecursion(node.Right)
}

// 中序遍历 的 非递归实现:
func inOrderIteration(node *TreeNode) {
	stack := []*TreeNode{}
	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%v ", cur.Val)
			node = cur.Right
		}
	}
}
