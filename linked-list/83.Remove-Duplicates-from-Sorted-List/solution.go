package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
