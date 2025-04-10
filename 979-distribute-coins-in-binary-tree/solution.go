package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var moves int

func distributeCoins(root *TreeNode) int {
	moves = 0
	dfs(root)
	return moves
}

func dfs(current *TreeNode) int {
	if current == nil {
		return 0
	}

	left := dfs(current.Left)
	right := dfs(current.Right)
	moves += abs(left) + abs(right)

	return current.Val - 1 + left + right

}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
