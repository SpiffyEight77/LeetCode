package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy

	for head.Next != nil {
		for head.Next != nil {
			if head.Next.Val == val {
				head.Next = head.Next.Next
			} else {
				head = head.Next
			}
		}
	}
	return dummy.Next
}
