package lcgo

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	var carry int
	current := &head
	const base = 10
	for p, q := l1, l2; p != nil || q != nil || carry != 0; {
		if p != nil {
			carry += p.Val
			p = p.Next
		}
		if q != nil {
			carry += q.Val
			q = q.Next
		}
		current.Next = &ListNode{Val: carry % base}
		carry = carry / base
		current = current.Next
	}
	return head.Next
}
