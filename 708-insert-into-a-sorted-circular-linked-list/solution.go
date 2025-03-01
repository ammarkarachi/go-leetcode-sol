package solution

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	node := &Node{
		Val: x,
	}
	if aNode == nil {

		node.Next = node
		return node
	}

	prev := aNode
	current := aNode.Next

	for {

		// case 1 typical
		if prev.Val <= x && current.Val >= x {
			break
		}

		//case 2
		if prev.Val > current.Val {
			if prev.Val <= x || current.Val >= x {
				break
			}
		}

		prev = current
		current = current.Next

		if prev == aNode {
			break
		}
	}

	prev.Next = node
	node.Next = current

	return aNode

}
