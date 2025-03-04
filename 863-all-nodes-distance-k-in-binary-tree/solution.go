package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	graph := make(map[int][]int)
	dfs(root, nil, &graph)

	queue := [][2]int{
		{target.Val, 0},
	}
	ans := []int{}
	visited := make(map[int]bool)
	visited[target.Val] = true
	for len(queue) > 0 {
		f := queue[0]
		queue = queue[1:]

		if f[1] == k {
			ans = append(ans, f[0])
			continue
		}

		n := graph[f[0]]

		for _, node := range n {
			if !visited[node] {
				visited[node] = true
				queue = append(queue, [2]int{node, f[1] + 1})
			}
		}

	}
	return ans
}

func dfs(node *TreeNode, parent *TreeNode, graph *map[int][]int) {
	g := *graph

	if node != nil && parent != nil {
		g[node.Val] = append(g[node.Val], parent.Val)
		g[parent.Val] = append(g[parent.Val], node.Val)
	}

	if node.Right != nil {
		dfs(node.Right, node, graph)
	}

	if node.Left != nil {
		dfs(node.Left, node, graph)
	}

}
