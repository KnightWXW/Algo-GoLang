package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 2, 5, 8, 3, 6, 9, 0, -1}
	fmt.Println(arr)
	selectSort(arr)
	fmt.Println(arr)
}

// Time: O(n ^ 2)
// Space: O(1)
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		arr[index], arr[i] = arr[i], arr[index]
	}
}
