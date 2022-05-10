package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 2, 5, 8, 3, 6, 9, 0, -1}
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}

// Time: O(n ^ 2)
// Space: O(1)
func bubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
