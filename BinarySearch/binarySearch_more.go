package main

import (
	"fmt"
)

func main() {
	// 在 有序数组 中 查找 第一个 大于 target 的 元素 的 索引位置
	arr := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch_more(arr, 7))
}

// Time: O(logn)
// Space: O(1)
func binarySearch_more(arr []int, target int) int {
	left, right := -1, len(arr)
	for left+1 != right {
		mid := left + (right-left)>>1
		if arr[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
