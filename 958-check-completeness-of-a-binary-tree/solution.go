package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	nilFound := false
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nil {
			nilFound = true
		} else {
			if nilFound {
				return false
			}

			q = append(q, node.Left)
			q = append(q, node.Right)
		}

	}

	return true
}
