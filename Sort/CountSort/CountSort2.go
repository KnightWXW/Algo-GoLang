package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 4, 7, 2, 5, 8, 3, 6, 9, 0, -1, 5, 8, 9, 7, 3, 8, 7}
	fmt.Println(arr)
	countSort2(arr)
	fmt.Println(arr)
}

// Time: O(n)
// Space: O(k)
// 该种 CountSort 无法保证排序的稳定性
func countSort2(arr []int) {
	minNum, maxNum := math.MaxInt64, math.MinInt64
	for i := 0; i < len(arr); i++ {
		minNum = min(arr[i], minNum)
		maxNum = max(arr[i], maxNum)
	}
	cnt := make([]int, maxNum-minNum+1)
	for i := 0; i < len(arr); i++ {
		cnt[arr[i]-minNum]++
	}
	index := 0
	for i := 0; i < len(cnt); i++ {
		for cnt[i] > 0 {
			arr[index] = minNum + i
			cnt[i]--
			index++
		}
	}
}
