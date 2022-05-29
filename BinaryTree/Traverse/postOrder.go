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
	fmt.Println("二叉树后序遍历(递归形式):")
	postOrderRecursion(tree.root)
	fmt.Println()
	fmt.Println("二叉树后序遍历(非递归形式):")
	postOrderIteration(tree.root)
}

// 后序遍历 的 递归实现:
func postOrderRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	postOrderRecursion(node.Left)
	postOrderRecursion(node.Right)
	fmt.Printf("%v ", node.Val)
}

// 后序遍历 的 非递归实现:
func postOrderIteration(node *TreeNode) {
	stack1, stack2 := []*TreeNode{node}, []*TreeNode{}
	for len(stack1) > 0 {
		cur := stack1[len(stack1)-1]
		stack2 = append(stack2, cur)
		stack1 = stack1[:len(stack1)-1]
		if cur.Left != nil {
			stack1 = append(stack1, cur.Left)
		}
		if cur.Right != nil {
			stack1 = append(stack1, cur.Right)
		}
	}

	for len(stack2) > 0 {
		cur := stack2[len(stack2)-1]
		fmt.Printf("%v ", cur.Val)
		stack2 = stack2[:len(stack2)-1]
	}
}
