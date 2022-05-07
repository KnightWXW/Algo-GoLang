package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 3, 9, -25, 52, 57, 63, 81}
	sort.Ints(arr)
	fmt.Println(binarySearch1(arr, -25))
}

func binarySearch1(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
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
