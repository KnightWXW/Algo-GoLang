package main

import "fmt"

func main() {
	arr := []int{1, 4, 3, 5, 5, 5, 2, 1, -1, 7, 0, 6, 7, 8, 9, 5}
	fmt.Println(arr)
	netherlandsFlag(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// Time: O(n)
// Space: O(1)
// 荷兰国旗问题：
// 给定一个数组arr，一个数pivot，请把小于pivot的数放在arr的左边，等于pivot的数放在arr的中间，
// 大于pivot的数放在arr的右边，不要求内部有序。
// 要求：时间复杂度O(N)，空间复杂度O(1)
func netherlandsFlag(arr []int, begin, end int) {
	left := begin - 1 // left：小于区域的右边界索引
	right := end      // right：大于区域的左边界索引
	i := begin
	for i < right {
		if arr[i] < arr[end] {
			arr[left+1], arr[i] = arr[i], arr[left+1]
			left++
			i++
		} else if arr[i] == arr[end] {
			i++
		} else {
			arr[right-1], arr[i] = arr[i], arr[right-1]
			right--
		}
	}
	arr[right], arr[end] = arr[end], arr[right]
}
