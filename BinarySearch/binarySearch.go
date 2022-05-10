package main

import (
	"fmt"
)

func main() {
	// 在 有序数组 中 查找 target 的 索引位置
	arr := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(arr, 7))
}

// Time: O(logn)
// Space: O(1)
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
