package main

func main() {

}

func isBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return process(root).Flag
}

func process(root *TreeNode) *InfoIsBST {
	if root == nil {
		return nil
	}
	leftInfo, rightInfo := process(root.Left), process(root.Right)
	if leftInfo.
}

type InfoIsBST struct {
	Flag bool
	MaxVal int
	MinVal int
}

func createInfoIsBST(b bool, maxVal, minVal int) *InfoIsBST {
	return &{b, maxVal, minVal}
}
