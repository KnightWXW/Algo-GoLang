package main

import "fmt"

// Heap 结构
type Heap struct {
	Arr   []int
	Size  int
	Limit int
}

func createHeap(limit int) *Heap {
	return &Heap{
		Arr:   make([]int, limit),
		Size:  0,
		Limit: limit,
	}
}

func main() {
	h := createHeap(10)
	fmt.Println(h.isEmpty()) //true
	h.push(7)
	h.push(3)
	h.push(5)
	fmt.Println(h.peek()) // 7
	h.push(9)
	h.push(1)
	fmt.Println(h.peek()) // 9
	fmt.Println(h.pop())  // 9
	fmt.Println(h.size()) // 4
	fmt.Println(h.peek()) // 7
	fmt.Print(h.Arr)      // [7 3 5 1 9 0 0 0 0 0]
}

// 判断 堆 是否为空
func (h *Heap) isEmpty() bool {
	return h.Size == 0
}

// 判断 堆 是否已满
func (h *Heap) isFull() bool {
	return h.Size == h.Limit
}

// 返回 堆的大小
func (h *Heap) size() int {
	return h.Size
}

// 返回 堆顶元素：
func (h *Heap) peek() int {
	return h.Arr[0]
}

// 向堆中 添加元素：
func (h *Heap) push(v int) {
	if h.isFull() {
		panic("The heap is full")
	}
	h.Arr[h.Size] = v
	h.heapInsert(h.Arr, h.Size)
	h.Size++
}

// 从堆中 弹出堆顶元素：
func (h *Heap) pop() int {
	if h.isEmpty() {
		panic("The heap is empty")
	}
	top := h.Arr[0]
	h.Size--
	h.Arr[0], h.Arr[h.Size] = h.Arr[h.Size], h.Arr[0]
	h.heapify(h.Arr, 0, h.Size)
	return top
}

// 向堆中 插入元素，插入的元素 不断往上探索，找到合适的位置
func (h *Heap) heapInsert(arr []int, i int) {
	for arr[i] > arr[(i-1)/2] {
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
		i = (i - 1) / 2
	}
}

// 调整 arr[i] 的位置，不断向下探索，形成堆结构
func (h *Heap) heapify(arr []int, i int, size int) {
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
