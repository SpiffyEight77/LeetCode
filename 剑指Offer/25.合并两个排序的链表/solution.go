package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return nil
	}

	dummy := &ListNode{0, nil}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}

	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}
	return dummy.Next
}
