package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{1, 4, 3, 5, 5, 5, 2, 1, -1, 7, 0, 6, 7, 8, 9, 5}
	fmt.Println(arr)
	quickSort1(arr)
	fmt.Println(arr)
}

func quickSort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	process1(arr, 0, len(arr)-1)
}

func process1(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	// 随机选择元素作为 pivot, 放到数组的末尾:
	i := rand.Intn(end-begin) + begin
	arr[i], arr[end] = arr[end], arr[i]
	indexs := partition1(arr, begin, end)
	process1(arr, begin, indexs[0]-1)
	process1(arr, indexs[1]+1, end)
}

func partition1(arr []int, begin, end int) []int {
	if begin > end {
		return []int{-1, -1}
	}
	if begin == end {
		return []int{begin, end}
	}
	left, right := begin-1, end
	index := begin
	for index < right {
		if arr[index] < arr[end] {
			arr[index], arr[left+1] = arr[left+1], arr[index]
			left++
			index++
		} else if arr[index] == arr[end] {
			index++
		} else {
			arr[index], arr[right-1] = arr[right-1], arr[index]
			right--
		}
	}
	arr[right], arr[end] = arr[end], arr[right]
	return []int{left + 1, right}
}
