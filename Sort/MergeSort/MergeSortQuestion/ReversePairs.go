package main

import "fmt"

func main() {
	arr := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(arr))
}

func reversePairs(arr []int) int {
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
	i, j, index := mid, right, len(nums)-1
	for i >= left && j >= mid+1 {

		if arr[i] > arr[j] {
			sum += j - mid
		}

		if arr[i] < arr[j] {
			nums[index] = arr[i]
			i--
		} else {
			nums[index] = arr[j]
			j--
		}
		index--
	}
	for i >= left {
		nums[index] = arr[i]
		index--
		i--
	}
	for j >= mid+1 {
		nums[index] = arr[j]
		index--
		j--
	}

	for i := 0; i < len(nums); i++ {
		arr[left+i] = nums[i]
	}
	return sum
}
