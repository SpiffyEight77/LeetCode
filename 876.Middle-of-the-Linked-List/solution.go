package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	p := head
	l := 0

	for p.Next != nil {
		l++
		p = p.Next
	}

	p = head
	l /= 2
	for l > 0 {
		l--
		p = p.Next
	}
	return p
}
