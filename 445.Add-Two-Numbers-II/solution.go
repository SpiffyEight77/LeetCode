package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	l1 = reverse(l1)
	l2 = reverse(l2)

	dummy := &ListNode{Val: 0}
	p := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{Val: carry % 10}
		carry /= 10
		p = p.Next
	}

	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev

		prev = head
		head = tmp
	}

	return prev
}
