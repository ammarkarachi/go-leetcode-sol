package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	var first *Node
	var last *Node
	var inOrder func(node *Node)

	inOrder = func(node *Node) {
		if node == nil {
			return
		}
		inOrder(node.Left)

		if last == nil {
			first = node
		} else {
			last.Right = node
			node.Left = last
		}

		last = node

		inOrder(node.Right)

	}
	inOrder(root)
	last.Right = first
	first.Left = last
	return first
}
