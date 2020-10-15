package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head

	var l int
	p := head
	for p != nil {
		l++
		p = p.Next
	}

	p = dummy
	l -= n
	for l > 0 {
		l--
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}
