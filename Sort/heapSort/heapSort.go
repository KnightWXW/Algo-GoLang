package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 2, 5, 8, 3, 6, 9, 0, -1}
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println(arr)
}

// O(nlogn)
func heapSort(arr []int) {
	size := len(arr)
	// 构建堆 方法一：
	// 利用 heapInsert 自上向下
	// Time: O(nlogn)
	for i := 0; i < size; i++ {
		heapInsert(arr, i)
	}
	/*// 构建堆 方法二：
	// 利用 heapify 自下向上
	// Time: O(n)
	for i := size - 1 ; i >= 0 ; i-- {
		heapify(arr, i, size)
	}*/

	// O(nlogn)
	size--
	arr[0], arr[size] = arr[size], arr[0]
	for size > 0 {
		heapify(arr, 0, size)
		size--
		arr[0], arr[size] = arr[size], arr[0]
	}
}

func heapInsert(arr []int, i int) {
	for arr[i] > arr[(i-1)/2] {
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
		i = (i - 1) / 2
	}
}

func heapify(arr []int, i int, size int) {
	left := 2*i + 1
	for left < size {
		large := left
		if left+1 < size && arr[left+1] > arr[left] {
			large = left + 1
		}
		if arr[i] > arr[large] {
			large = i
		}
		if i == large {
			break
		}
		arr[i], arr[large] = arr[large], arr[i]
		i = large
		left = 2*i + 1
	}
}
