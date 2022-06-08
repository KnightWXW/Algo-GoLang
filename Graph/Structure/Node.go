package main

type Node struct {
	InDegree  int     `入度`
	OutDegree int     `出度`
	NextNode  []*Node `相邻节点集合`
	NextEdge  []*Edge `相邻边集合`
}

func createNode() *Node {
	return &Node{0, 0, []*Node{}, []*Edge{}}
}
