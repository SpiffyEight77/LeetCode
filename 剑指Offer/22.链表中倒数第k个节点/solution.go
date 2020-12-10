package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	dummy := &ListNode{0, head}
	head = dummy

	for i := 0; i < k; i++ {
		head = head.Next
	}

	for head != nil {
		head = head.Next
		dummy = dummy.Next
	}
	return dummy
}
