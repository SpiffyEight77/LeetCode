package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	head = dummy

	for head.Next != nil && head.Next.Next != nil {
		node1 := head.Next
		node2 := head.Next.Next
		head.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		head = head.Next.Next
	}
	return dummy.Next
}
