package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	count := 1
	curr := head
	for curr != nil && curr.Next != nil {
		curr = curr.Next
		count++
	}

	if count == 1 {
		return head
	}

	curr.Next = head
	newTail := head
	for i := 1; i < count-k%count; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next

	newTail.Next = nil

	return newHead

}
