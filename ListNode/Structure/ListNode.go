package main

import "fmt"

type ListNode struct {
	Val  interface{} `数据域：当前节点存储的数值`
	Next *ListNode   `指针域：当前节点的下一节点`
}

type LinkedList struct {
	Head *ListNode `链表的首结点`
	Tail *ListNode `链表的尾结点`
}

// 创建链表节点：
func CreateListNode(val int) *ListNode {
	return &ListNode{val, nil}
}

// 根据数组创建链表
func (this *LinkedList) CreateLinkedListFromArray(arr []int) {
	if len(arr) == 0 {
		return
	}
	node := &ListNode{arr[0], nil}
	this.Head = node
	cur := node
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{arr[i], nil}
		cur = cur.Next
	}
	this.Tail = cur
}

func PrintListNode(node *ListNode) {
	arr := []interface{}{}
	cur := node
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	fmt.Println(arr)
}

// 判断 链表是否 为空：
func (this *LinkedList) IsEmpty() bool {
	return this.Head == nil
}

// 判断 是否 最后一个节点:
func (this *LinkedList) IsLastNode(node *ListNode) bool {
	return node.Next == nil
}

// 返回链表的长度：
func (this *LinkedList) SizeOfListNode() int {
	if this.Head == nil {
		return 0
	}
	node := this.Head
	size := 1
	for node.Next != nil {
		node = node.Next
		size++
	}
	return size
}

// 查找 后继节点:
func (this *LinkedList) FindNextNode(node *ListNode) *ListNode {
	return node.Next
}

// 在 链表头部 添加数据：
func (this *LinkedList) Add(val int) {
	node := &ListNode{val, nil}
	node.Next = this.Head
	this.Head = node
}

// 在 链表尾部 追加数据：
func (this *LinkedList) Append(val int) {
	node := &ListNode{val, nil}
	this.Tail.Next = node
	this.Tail = this.Tail.Next
}

// 在指定位置 插入数据
func (this *LinkedList) Insert(index int, val int) {
	node := &ListNode{val, nil}
	if index < 0 {
		fmt.Println("Index out of range.")
	}
	cur := this.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	tem := cur.Next
	cur.Next = node
	node.Next = tem
}

// 删除 指定index位置 的 数据：
func (this *LinkedList) Delete(index int) {
	if this.Head == nil {
		return
	}
	if index < 0 {
		fmt.Println("Index out of range.")
		return
	} else if index == 0 {
		this.Head = this.Head.Next
		return
	}

	cur := this.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var linkedList LinkedList
	fmt.Println(linkedList.IsEmpty())
	linkedList.CreateLinkedListFromArray(arr)
	PrintListNode(linkedList.Head)
	fmt.Println(linkedList.SizeOfListNode())
	// 此时为 1->2->3->4->5->6->7
	linkedList.Add(0)
	PrintListNode(linkedList.Head)
	// 此时为 0->1->2->3->4->5->6->7
	linkedList.Append(8)
	PrintListNode(linkedList.Head)
	// 此时为 0->1->2->3->4->5->6->7->8
	linkedList.Insert(3, -1)
	PrintListNode(linkedList.Head)
	// 此时为 0->1->2->3->4->-1->5->6->7->8
	linkedList.Delete(3)
	PrintListNode(linkedList.Head)
	// 此时为 0->1->2->3->4->5->6->7->8
	linkedList.Delete(0)
	PrintListNode(linkedList.Head)
	// 此时为 0->1->2->3->4->5->6->7->8
}
