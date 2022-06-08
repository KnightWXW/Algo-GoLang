package main

import "fmt"

type TrieofArray struct {
	Pass    int            `当前节点被通过几次`
	End     int            `当前节点多少次被当作结尾节点`
	NextArr []*TrieofArray `下一层的后续节点数组`
}

func createTrieofArray() *TrieofArray {
	return &TrieofArray{0, 0, make([]*TrieofArray, 26)}
}

// 在 前缀树 中插入 字符串word：
func (this *TrieofArray) InsertofArray(word string) {
	if word == "" {
		return
	}
	node := this
	node.Pass++
	path := 0
	for i := 0; i < len(word); i++ {
		path = int(word[i] - 'a')
		if node.NextArr[path] == nil {
			node.NextArr[path] = createTrieofArray()
		}
		node = node.NextArr[path]
		node.Pass++
	}
	node.End++
}

// 判断字符串 word 之前在前缀树中 出现过 多少次：
func (this *TrieofArray) SearchofArray(word string) int {
	if word == "" {
		return 0
	}
	node := this
	path := 0
	for i := 0; i < len(word); i++ {
		path = int(word[i] - 'a')
		if node.NextArr[path] == nil {
			return 0
		}
		node = node.NextArr[path]
	}
	return node.End
}

// 删除 在前缀树中 的 字符串 word：
func (this *TrieofArray) DeleteofArray(word string) {
	if this.SearchofArray(word) != 0 {
		node := this
		node.Pass--
		path := 0
		for i := 0; i < len(word); i++ {
			path = int(word[i] - 'a')
			if node.NextArr[path].Pass == 1 {
				node.NextArr[path] = nil
				return
			}
			node = node.NextArr[path]
		}
		node.End--
	}
}

// 在前缀树中，多少个以 pre字符串 作为前缀：
func (this *TrieofArray) prefixNumberofArray(pre string) int {
	if pre == "" {
		return 0
	}
	node := this
	path := 0
	for i := 0; i < len(pre); i++ {
		path = int(pre[i] - 'a')
		if node.NextArr[path] == nil {
			return 0
		}
		node = node.NextArr[path]
	}
	return node.Pass
}

func main() {
	s1, s2, s3 := "apple", "apples", "apple"
	root := createTrieofArray()
	root.InsertofArray(s1)
	root.InsertofArray(s2)
	root.InsertofArray(s3)
	fmt.Println(root.SearchofArray("appl"))
	fmt.Println(root.SearchofArray("apple"))
	fmt.Println(root.prefixNumberofArray("app"))
	root.DeleteofArray(s1)
	fmt.Println(root.SearchofArray("appl"))
	fmt.Println(root.SearchofArray("apple"))
	fmt.Println(root.prefixNumberofArray("app"))
}
