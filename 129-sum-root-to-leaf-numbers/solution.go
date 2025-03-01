package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeSum struct {
	Sum  int
	Node *TreeNode
}

func sumNumbers(root *TreeNode) int {

	stack := []*TreeNodeSum{{
		Node: root,
		Sum:  0,
	}}
	sum := 0
	for len(stack) > 0 {
		treeNodeSum := stack[0]
		stack = stack[1:]
		currentSum := treeNodeSum.Sum
		if treeNodeSum.Node.Left == nil && treeNodeSum.Node.Right == nil {
			sum += currentSum*10 + treeNodeSum.Node.Val
		}

		if treeNodeSum.Node.Right != nil {
			stack = append([]*TreeNodeSum{{
				Node: treeNodeSum.Node.Right,
				Sum:  currentSum*10 + treeNodeSum.Node.Val,
			}}, stack...)
		}

		if treeNodeSum.Node.Left != nil {
			stack = append([]*TreeNodeSum{{
				Node: treeNodeSum.Node.Left,
				Sum:  currentSum*10 + treeNodeSum.Node.Val,
			}}, stack...)
		}
	}
	return sum
}
