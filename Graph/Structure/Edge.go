package main

type Edge struct {
	Weight int   `权重`
	From   *Node `边的进入节点`
	To     *Node `边的出去节点`
}

func createEdge() *Edge {
	return &Edge{0, nil, nil}
}
