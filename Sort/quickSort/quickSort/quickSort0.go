package main

import "fmt"

func main() {
	arr := []int{1, 4, 3, 5, 5, 5, 2, 1, -1, 7, 0, 6, 7, 8, 9, 5}
	fmt.Println(arr)
	quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	process0(arr, 0, len(arr)-1)
}

func process0(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	indexs := partition0(arr, begin, end)
	process0(arr, begin, indexs[0]-1)
	process0(arr, indexs[1]+1, end)
}

func partition0(arr []int, begin, end int) []int {
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
