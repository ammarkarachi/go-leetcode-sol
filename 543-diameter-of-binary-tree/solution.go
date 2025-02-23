package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	depth(root, &max)
	return max
}

func depth(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left, max)
	right := depth(node.Right, max)
	*max = maxNum(*max, left+right)
	return maxNum(left, right) + 1
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
