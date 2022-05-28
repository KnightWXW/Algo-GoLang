package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 4, 7, 2, 5, 8, 3, 6, 9, 0, -1, 5, 8, 9, 7, 3, 8, 7}
	fmt.Println(arr)
	countSort(arr)
	fmt.Println(arr)
}

// Time: O(n)
// Space: O(n)
// 该种 countSort 可以保证排序的稳定性
func countSort(arr []int) {
	minNum, maxNum := math.MaxInt64, math.MinInt64
	for i := 0; i < len(arr); i++ {
		minNum = min(arr[i], minNum)
		maxNum = max(arr[i], maxNum)
	}
	cnt := make([]int, maxNum-minNum+1)
	for i := 0; i < len(arr); i++ {
		cnt[arr[i]-minNum]++
	}
	for i := 1; i < len(cnt); i++ {
		cnt[i] += cnt[i-1]
	}
	ans := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i] - minNum
		index := cnt[v] - 1
		ans[index] = arr[i]
		cnt[v]--
	}
	copy(arr, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
