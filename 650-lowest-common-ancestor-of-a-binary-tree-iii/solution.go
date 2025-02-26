package solution

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {

	set := make(map[*Node]bool)

	stack := make([]*Node, 0)

	stack = append(stack, q.Parent, p.Parent)

	set[q.Parent] = true
	set[p.Parent] = true

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		_, ok := set[node]
		if ok {
			return node
		} else {
			set[node.Parent] = true
		}
	}

	return nil
}
