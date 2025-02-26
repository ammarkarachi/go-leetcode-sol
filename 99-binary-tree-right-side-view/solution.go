package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeLevel struct {
	Node  *TreeNode
	Level int
}

func rightSideView(root *TreeNode) []int {
	lastSeen := make(map[int]int)
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNodeLevel, 0)
	stack = append(stack, &TreeNodeLevel{
		Node:  root,
		Level: 0,
	})

	for len(stack) > 0 {
		nodeLevel := pop(&stack)
		lastSeen[nodeLevel.Level] = nodeLevel.Node.Val

		if nodeLevel.Node.Right != nil {
			push(&stack, &TreeNodeLevel{
				Node:  nodeLevel.Node.Right,
				Level: nodeLevel.Level + 1,
			})
		}

		if nodeLevel.Node.Left != nil {
			push(&stack, &TreeNodeLevel{
				Node:  nodeLevel.Node.Left,
				Level: nodeLevel.Level + 1,
			})
		}

	}

	ans := make([]int, len(lastSeen))
	for k, v := range lastSeen {
		ans[k] = v
	}
	return ans

}

func push[T any](stack *[]T, elem T) {
	*stack = append(*stack, elem)
}

func pop[T any](stack *[]T) T {
	if len(*stack) == 0 {
		var zero T
		return zero
	}
	elem := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return elem
}
