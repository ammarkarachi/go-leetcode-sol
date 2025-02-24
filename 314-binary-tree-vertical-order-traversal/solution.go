package solution

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeLocation struct {
	Node     *TreeNode
	Location int
}

func verticalOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNodeLocation, 0)
	queue = append(queue, &TreeNodeLocation{
		Node:     root,
		Location: 0,
	})
	if root == nil {
		return make([][]int, 0)
	}
	nodeMap := make(map[int][]*TreeNode)
	min := math.MaxInt
	max := math.MinInt
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if min > node.Location {
			min = node.Location
		}
		if max < node.Location {
			max = node.Location
		}

		if _, ok := nodeMap[node.Location]; !ok {
			nodeMap[node.Location] = make([]*TreeNode, 0)
		}
		nodeMap[node.Location] = append(nodeMap[node.Location], node.Node)

		if node.Node.Left != nil {
			queue = append(queue, &TreeNodeLocation{
				Node:     node.Node.Left,
				Location: node.Location - 1,
			})
		}

		if node.Node.Right != nil {
			queue = append(queue, &TreeNodeLocation{
				Node:     node.Node.Right,
				Location: node.Location + 1,
			})
		}

	}

	answer := make([][]int, max-min+1)
	for k, v := range nodeMap {
		idx := k - min
		answer[idx] = make([]int, len(v))
		for i, node := range v {
			answer[idx][i] = node.Val
		}
	}
	return answer
}
