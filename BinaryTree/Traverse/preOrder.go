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
	fmt.Println("二叉树前序遍历(递归形式):")
	preOrderRecursion(tree.root)
	fmt.Println()
	fmt.Println("二叉树前序遍历(非递归形式):")
	preOrderIteration(tree.root)
}

// 前序遍历 的 递归实现:
func preOrderRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.Val)
	preOrderRecursion(node.Left)
	preOrderRecursion(node.Right)
}

// 前序遍历 的 非递归实现:
func preOrderIteration(node *TreeNode) {
	stack := []*TreeNode{node}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		fmt.Printf("%v ", cur.Val)
		stack = stack[:len(stack)-1]
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
}
