package main

import "fmt"

type UnionFindofArray struct {
	Parents []int `存储当前节点的父亲节点`
	SizeMap []int `存储以当前节点为祖先节点的节点集合的数量大小`
	Help    []int `辅助数组`
	Sets    int   `并查集含有的集合个数`
}

func createUnionFindofArray(n int) *UnionFindofArray {
	parent, sizeMap, help := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		sizeMap[i] = 1
	}
	return &UnionFindofArray{
		parent, sizeMap, help, n,
	}
}

// 合并当前两个节点的集合：
func (this *UnionFindofArray) unionofArray(a, b int) {
	headA, headB := this.findofArray(a), this.findofArray(b)
	if headA != headB {
		if this.SizeMap[headA] >= this.SizeMap[headB] {
			this.SizeMap[headA] += this.SizeMap[headB]
			this.Parents[headB] = headA
		} else {
			this.SizeMap[headB] += this.SizeMap[headA]
			this.Parents[headA] = headB
		}
		this.Sets--
	}
}

// 查找当前节点的祖先节点：
func (this *UnionFindofArray) findofArray(a int) int {
	for a != this.Parents[a] {
		this.Help = append(this.Help, a)
		a = this.Parents[a]
	}
	// 路径压缩：
	for len(this.Help) != 0 {
		cur := this.Help[len(this.Help)-1]
		this.Help = this.Help[:len(this.Help)-1]
		this.Parents[cur] = a
	}
	return a
}

// 判断当前两个节点是否属于同一集合：
func (this *UnionFindofArray) isSameSetofArray(a, b int) bool {
	return this.findofArray(a) == this.findofArray(b)
}

// 返回并查集中集合的数目：
func (this *UnionFindofArray) sizeofArray() int {
	return this.Sets
}

func main() {
	arr := [][]int{{1, 1, 0, 0}, {1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}
	uf := createUnionFindofArray(len(arr))
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr[0]); j++ {
			if arr[i][j] == 1 {
				uf.unionofArray(i, j)
			}
		}
	}
	fmt.Println(uf.sizeofArray())
}
