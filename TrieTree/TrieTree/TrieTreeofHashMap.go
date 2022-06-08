package main

import "fmt"

type TrieofHashMap struct {
	Pass    int                    `当前节点被通过几次`
	End     int                    `当前节点多少次被当作结尾节点`
	NextMap map[int]*TrieofHashMap `下一层的后续节点数组`
}

func createTrieofHashMap() *TrieofHashMap {
	return &TrieofHashMap{0, 0, map[int]*TrieofHashMap{}}
}

// 在 前缀树 中插入 字符串word：
func (this *TrieofHashMap) Insert(word string) {
	if word == "" {
		return
	}
	node := this
	node.Pass++
	path := 0
	for i := 0; i < len(word); i++ {
		path = int(word[i] - 'a')
		if _, ok := node.NextMap[path]; !ok {
			node.NextMap[path] = createTrieofHashMap()
		}
		node = node.NextMap[path]
		node.Pass++
	}
	node.End++
}

// 判断字符串 word 之前在前缀树中 出现过 多少次：
func (this *TrieofHashMap) Search(word string) int {
	if word == "" {
		return 0
	}
	node := this
	path := 0
	for i := 0; i < len(word); i++ {
		path = int(word[i] - 'a')
		if _, ok := node.NextMap[path]; !ok {
			return 0
		}
		node = node.NextMap[path]
	}
	return node.End
}

// 删除 在前缀树中 的 字符串 word：
func (this *TrieofHashMap) Delete(word string) {
	if this.Search(word) != 0 {
		node := this
		node.Pass--
		path := 0
		for i := 0; i < len(word); i++ {
			path = int(word[i] - 'a')
			if _, ok := node.NextMap[path]; !ok {
				node.NextMap[path] = nil
				return
			}
			node = node.NextMap[path]
		}
		node.End--
	}
}

// 在前缀树中，多少个以 pre字符串 作为前缀：
func (this *TrieofHashMap) prefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	node := this
	path := 0
	for i := 0; i < len(pre); i++ {
		path = int(pre[i] - 'a')
		if _, ok := node.NextMap[path]; !ok {
			return 0
		}
		node = node.NextMap[path]
	}
	return node.Pass
}

func main() {
	s1, s2, s3 := "apple", "apples", "apple"
	root := createTrieofHashMap()
	root.Insert(s1)
	root.Insert(s2)
	root.Insert(s3)
	fmt.Println(root.Search("appl"))
	fmt.Println(root.Search("apple"))
	fmt.Println(root.prefixNumber("app"))
	root.Delete(s1)
	fmt.Println(root.Search("appl"))
	fmt.Println(root.Search("apple"))
	fmt.Println(root.prefixNumber("app"))
}
