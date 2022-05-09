package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 2, 5, 8, 3, 6, 9, 0, -1}
	fmt.Println(arr)
	mergeSort_(arr)
	fmt.Println(arr)
}

// 归并排序 的 非递归版本
// Time: O(nlogn)
// Space: O(n)
func mergeSort_(arr []int) {
	n := len(arr)
	step := 1
	for step < n {
		left := 0
		for left < n {
			mid := left + step - 1
			if mid >= n {
				break
			}
			right := min(mid+step, n-1)
			merge(arr, left, mid, right)
			left = right + 1
		}
		if step > n/2 {
			break
		}
		step <<= 1
	}
}

func merge(arr []int, left, mid, right int) {
	nums := make([]int, right-left+1)
	i, j, index := left, mid+1, 0
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			nums[index] = arr[i]
			i++
		} else {
			nums[index] = arr[j]
			j++
		}
		index++
	}
	for i <= mid {
		nums[index] = arr[i]
		index++
		i++
	}
	for j <= right {
		nums[index] = arr[j]
		index++
		j++
	}

	for i := 0; i < len(nums); i++ {
		arr[left+i] = nums[i]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
