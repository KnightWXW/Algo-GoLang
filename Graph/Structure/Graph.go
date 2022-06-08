package main

type Graph struct {
	NodeMap map[int]*Node  `节点的集合`
	EdgeMap map[*Edge]bool `边的集合`
}

func createGraph() *Graph {
	return &Graph{map[int]*Node{}, map[*Edge]bool{}}
}
