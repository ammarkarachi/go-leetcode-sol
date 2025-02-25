package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	sum += rangeSumBST(root.Left, low, high)
	sum += rangeSumBST(root.Right, low, high)
	return sum
}
