package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}
	head = dummy

	for head != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
	return dummy.Next
}
