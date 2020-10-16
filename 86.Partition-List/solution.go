package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	headDummy.Next = head
	head = headDummy
	tail := tailDummy

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			t := head.Next
			tail.Next = t
			tail = tail.Next
			head.Next = head.Next.Next
		}
	}

	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}
