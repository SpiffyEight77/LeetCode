package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		node1 := prev.Next
		node2 := prev.Next.Next
		prev.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		prev = node1
	}
	return dummy.Next
}
