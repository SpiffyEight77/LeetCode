package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	l := 0

	for head.Next != nil {
		l++
		head = head.Next
	}
	head.Next = dummy.Next

	k %= l
	for i := 0; i < l-k; i++ {
		head = head.Next
	}

	dummy.Next = head.Next
	head.Next = nil

	return dummy.Next
}
