package tool

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func (this *TreeNode) Create(val int) {
	this.val = val
}
