package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	N := n
	first := dummy
	second := head

	for N > 0 {
		second = second.Next
		N--
	}

	for second != nil {
		first = first.Next
		second = second.Next
	}

	first.Next = first.Next.Next

	return dummy.Next

}
