package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenhead := head.Next
	for odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenhead
	return head
}

func oddEvenList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := &ListNode{Val: 0}
	even := &ListNode{Val: 0}
	p, q := odd, even
	togger := true

	for head != nil {
		if togger {
			p.Next = &ListNode{Val: head.Val}
			p = p.Next
		} else {
			q.Next = &ListNode{Val: head.Val}
			q = q.Next
		}
		togger = !togger
		head = head.Next
	}
	p.Next = even.Next

	return odd.Next
}
