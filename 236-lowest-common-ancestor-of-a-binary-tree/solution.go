package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	parent := make(map[*TreeNode]*TreeNode)

	parent[root] = nil
	stack = append(stack, root)
	_, okQ := parent[q]
	_, okP := parent[p]
	for !okQ || !okP {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			parent[node] = node.Left
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			parent[node] = node.Right
			stack = append(stack, node.Right)
		}

		_, okP = parent[p]
		_, okQ = parent[q]

	}

	ancestors := make(map[*TreeNode]bool)

	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}

	for _, ok := ancestors[q]; ok; {
		node, _ := parent[q]
		q = node
	}
	return q
}
