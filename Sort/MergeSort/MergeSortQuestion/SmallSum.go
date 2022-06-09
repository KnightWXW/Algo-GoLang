package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 2, 5, 8, 3, 6, 9, 0, -1}
	fmt.Println(smallSum(arr))
}

func smallSum(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)>>1
	return process(arr, left, mid) +
		process(arr, mid+1, right) +
		merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) int {
	sum := 0
	nums := make([]int, right-left+1)
	i, j, index := left, mid+1, 0
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			sum += (right - j + 1) * arr[i]
		}
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
	return sum
}
