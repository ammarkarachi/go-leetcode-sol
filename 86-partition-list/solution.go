package solutio

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy1 := &ListNode{
		Val: -1,
	}

	dummy2 := &ListNode{
		Val: -2,
	}

	newHead := dummy1
	newTail := dummy2
	curr := head

	for curr != nil {
		if curr.Val < x {
			newHead.Next = curr
			newHead = newHead.Next
		} else {
			newTail.Next = curr
			newTail = newTail.Next
		}
		curr = curr.Next
	}

	newTail.Next = nil
	newHead.Next = dummy2.Next
	return dummy1.Next
}
