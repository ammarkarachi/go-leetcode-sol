package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{
		stack: make([]*TreeNode, 0),
	}
	iterator.leftMost(root)
	return iterator
}

func (this *BSTIterator) leftMost(node *TreeNode) {
	n := node
	for n != nil {
		this.stack = append(this.stack, n)
		n = n.Left
	}
}

func (this *BSTIterator) Next() int {
	length := len(this.stack)
	top := this.stack[length-1]
	this.stack = this.stack[:length-1]
	if top.Right != nil {
		this.leftMost(top.Right)
	}
	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
