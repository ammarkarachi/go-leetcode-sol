package solution

import (
	"math"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeLocation struct {
	Node  *TreeNode
	Width int
	Depth int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNodeLocation{{Node: root, Width: 0, Depth: 0}}
	treeMap := make(map[[2]int][]int)
	minWidth := math.MaxInt
	maxWidth := math.MinInt
	maxDepth := math.MinInt

	for len(stack) > 0 {
		node := pop(&stack)
		minWidth = min(minWidth, node.Width)
		maxWidth = max(maxWidth, node.Width)
		maxDepth = max(maxDepth, node.Depth)
		key := [2]int{node.Width, node.Depth}
		arr, exists := treeMap[key]
		if exists {
			treeMap[key] = append(arr, node.Node.Val)
		} else {
			treeMap[key] = []int{node.Node.Val}
		}
		nextDepth := node.Depth + 1
		if node.Node.Right != nil {
			width := node.Width + 1
			push(&stack, &TreeNodeLocation{
				Node:  node.Node.Right,
				Width: width,
				Depth: nextDepth,
			})

		}

		if node.Node.Left != nil {
			width := node.Width - 1
			push(&stack, &TreeNodeLocation{
				Node:  node.Node.Left,
				Width: width,
				Depth: nextDepth,
			})
		}
	}
	ans := make([][]int, maxWidth-minWidth+1)

	for i := minWidth; i <= maxWidth; i++ {
		idx := i - minWidth
		columnVal := []int{}
		for j := 0; j <= maxDepth; j++ {
			val, ok := treeMap[[2]int{i, j}]
			if ok {
				slices.Sort(val)
				columnVal = append(columnVal, val...)
			}

		}
		ans[idx] = columnVal
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
