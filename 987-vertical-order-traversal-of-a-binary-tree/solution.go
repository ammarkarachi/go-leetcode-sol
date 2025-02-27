package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeLocation struct {
	Node  *TreeNode
	Width int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNodeLocation{{Node: root, Width: 0}}
	treeMap := make(map[int][]int)
	minWidth := math.MaxInt
	maxWidth := math.MinInt
	treeMap[0] = []int{root.Val}

	for len(stack) > 0 {
		node := pop(&stack)
		minWidth = min(minWidth, node.Width)
		maxWidth = max(maxWidth, node.Width)

		arr, exists := treeMap[node.Width]
		if exists {
			treeMap[node.Width] = append(arr, node.Node.Val)
		} else {
			treeMap[node.Width] = []int{node.Node.Val}
		}

		if node.Node.Right != nil {
			width := node.Width + 1
			push(&stack, &TreeNodeLocation{
				Node:  node.Node.Right,
				Width: width,
			})

		}

		if node.Node.Left != nil {
			width := node.Width - 1
			push(&stack, &TreeNodeLocation{
				Node:  node.Node.Left,
				Width: width,
			})
		}
	}
	ans := make([][]int, maxWidth-minWidth+1)
	for k, v := range treeMap {
		idx := k - minWidth
		ans[idx] = v
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func push(stack *[]*TreeNodeLocation, item *TreeNodeLocation) {
	(*stack) = append([]*TreeNodeLocation{item}, (*stack)...)
}

func pop(stack *[]*TreeNodeLocation) *TreeNodeLocation {
	s := (*stack)
	item := s[0]
	(*stack) = s[1:]
	return item
}
