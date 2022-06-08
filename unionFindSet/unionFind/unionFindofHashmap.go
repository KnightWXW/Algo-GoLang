package main

import "fmt"

type Node struct {
	Value int `以value为值的节点`
}

func createNode(v int) *Node {
	return &Node{v}
}

type UnionFindofHashMap struct {
	Nodes   map[int]*Node   `存储当前节点`
	Parents map[*Node]*Node `存储当前节点的父亲节点`
	SizeMap map[*Node]int   `存储以当前节点为祖先节点的节点集合的数量大小`
}

func createUnionFindofHashMap(arr []int) *UnionFindofHashMap {
	Nodes := make(map[int]*Node)
	Parents := make(map[*Node]*Node)
	SizeMap := make(map[*Node]int)
	for _, v := range arr {
		cur := createNode(v)
		Nodes[v] = cur
		Parents[cur] = cur
		SizeMap[cur] = 1
	}
	return &UnionFindofHashMap{Nodes, Parents, SizeMap}
}

// 合并当前两个节点的集合：
func (this *UnionFindofHashMap) unionofHashMap(curA, curB *Node) {
	headA, headB := this.findofHashMap(curA), this.findofHashMap(curB)
	if headA != headB {
		sizeA, sizeB := this.SizeMap[headA], this.SizeMap[headB]
		big, small := headA, headB
		if sizeA <= sizeB {
			big, small = headB, headA
		}
		this.Parents[small] = big
		this.SizeMap[big] = sizeA + sizeB
		delete(this.SizeMap, small)
	}
}

// 查找当前节点的祖先节点：
func (this *UnionFindofHashMap) findofHashMap(cur *Node) *Node {
	stack := []*Node{}
	for cur != this.Parents[cur] {
		stack = append(stack, cur)
		cur = this.Parents[cur]
	}
	// 路径压缩：
	for len(stack) != 0 {
		tem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		this.Parents[tem] = cur
	}
	return cur
}

// 判断当前两个节点是否属于同一集合：
func (this *UnionFindofHashMap) isSameSetofHashMap(a int, b int) bool {
	return this.findofHashMap(this.Nodes[a]) == this.findofHashMap(this.Nodes[b])
}

// 返回并查集中集合的数目：
func (this *UnionFindofHashMap) sizeofHashMap() int {
	return len(this.SizeMap)
}

func main() {
	arr := [][]int{{1, 1, 0, 0}, {1, 1, 0, 0}, {0, 0, 1, 1}, {0, 0, 1, 1}}
	a := []int{0, 1, 2, 3}
	uf := createUnionFindofHashMap(a)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr[0]); j++ {
			if arr[i][j] == 1 {
				uf.unionofHashMap(uf.Nodes[i], uf.Nodes[j])
			}
		}
	}
	fmt.Println(uf.sizeofHashMap())
}
