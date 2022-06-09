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

	fmt.Println("二叉树的层序遍历:")
	ans := levelOrderTraverse(tree.root)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}

// 层序遍历:
func levelOrderTraverse(node *TreeNode) [][]int {
	if node == nil {
		return [][]int{}
	}
	ans := [][]int{}
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		size := len(queue)
		list := []int{}
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			list = append(list, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ans = append(ans, list)
	}
	return ans
}
