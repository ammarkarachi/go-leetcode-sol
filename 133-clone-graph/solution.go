package solution

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	visited := make(map[*Node]*Node)
	stack := []*Node{}

	stack = append(stack, node)

	visited[node] = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}

	for len(stack) > 0 {
		stackNode := stack[0]
		stack = stack[1:]
		copyNode := visited[stackNode]

		for i, neighbor := range stackNode.Neighbors {
			_, exists := visited[neighbor]
			if !exists {
				visited[neighbor] = &Node{
					Val:       neighbor.Val,
					Neighbors: make([]*Node, len(neighbor.Neighbors)),
				}
				stack = append([]*Node{neighbor}, stack...)
			}
			copyNode.Neighbors[i] = visited[neighbor]
		}
	}

	return visited[node]
}
