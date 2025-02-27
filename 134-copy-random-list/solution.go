package solution

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	copyMap := make(map[*Node]*Node)
	queue := []*Node{head}
	if head == nil {
		return head
	}
	// create the copy
	newNode := &Node{
		Val: head.Val,
	}

	// add the copy to the map
	copyMap[head] = newNode

	for len(queue) > 0 {
		node := poll(&queue)
		copyNode := copyMap[node]

		// check if next is there
		if node.Next != nil {
			// get if the copied node
			nextNode, ok := copyMap[node.Next]
			// if not create one and assign
			if !ok {
				prepend(&queue, node.Next)
				copyNextNode := &Node{
					Val: node.Next.Val,
				}
				copyMap[node.Next] = copyNextNode
				copyNode.Next = copyNextNode
			} else {
				// if there is assign since we already assign it's children in an earlier stage
				copyNode.Next = nextNode
			}

		}

		if node.Random != nil {
			randomNode, ok := copyMap[node.Random]
			if !ok {
				prepend(&queue, node.Random)
				copyRandomNode := &Node{
					Val: node.Random.Val,
				}
				copyMap[node.Random] = copyRandomNode
				copyNode.Random = copyRandomNode
			} else {
				copyNode.Random = randomNode
			}
		}

	}
	return copyMap[head]
}

func prepend(queue *[]*Node, item *Node) {
	(*queue) = append([]*Node{item}, (*queue)...)
}

func poll(queue *[]*Node) *Node {
	n := len(*queue)
	q := *queue
	item := q[n-1]
	(*queue) = q[:n-1]
	return item
}
