package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{15, 154, 257, 42, 55, 80, 43, 76, 59, 30, 31, 52, 0}
	fmt.Println(arr)
	radixSort(arr)
	fmt.Println(arr)
}

// Time: O(d(n + r))    r为基数, d为位数
// Space: O(n + r)
func radixSort(arr []int) {
	digit := cntMaxofLen(arr)
	radixSortProcess(arr, 0, len(arr)-1, digit)
}

func radixSortProcess(arr []int, left, right, digit int) {
	const RADIX = 10
	help := make([]int, right-left+1)
	for k := 1; k <= digit; k++ {
		// 开辟 辅助数组cnt：
		cnt := make([]int, RADIX)
		// 将 arr 按 arr[i]的 第k位数字 存储到cnt数组中：
		for i := left; i <= right; i++ {
			tem := getNumofDigit(arr[i], k)
			cnt[tem]++
		}
		// cnt数组 转换成 累加和：
		for i := 1; i < len(cnt); i++ {
			cnt[i] += cnt[i-1]
		}
		// 根据 cnt数组 确定 元素顺序
		for i := right; i >= left; i-- {
			tem := getNumofDigit(arr[i], k)
			help[cnt[tem]-1] = arr[i]
			cnt[tem]--
		}

		// 将 help数组 的元素 转移至 arr数组
		j := 0
		for i := left; i <= right; i++ {
			arr[i] = help[j]
			j++
		}
	}
}

// 获得 最大元素 的 数位的个数：
func cntMaxofLen(arr []int) int {
	maxNum := math.MinInt64
	for i := 0; i < len(arr); i++ {
		if maxNum > arr[i] {
			maxNum = arr[i]
		}
	}
	cnt := 0
	for maxNum != 0 {
		maxNum /= 10
		cnt++
	}
	return cnt
}

// 获得 num 的 第 k 位数字：
func getNumofDigit(num, k int) int {
	return (num / int(math.Pow(10, float64(k-1))) % 10)
}
