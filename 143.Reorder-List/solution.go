package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow

	p := mid.Next
	var prev *ListNode
	for p != nil {
		temp := p.Next
		p.Next = prev
		prev = p
		p = temp
	}

	mid.Next = nil
	l1 := head
	l2 := prev
	dummy := &ListNode{Val: 0}
	dummyHead := dummy
	toggle := true
	for l1 != nil && l2 != nil {
		if toggle {
			dummyHead.Next = l1
			l1 = l1.Next
		} else {
			dummyHead.Next = l2
			l2 = l2.Next
		}
		toggle = !toggle
		dummyHead = dummyHead.Next
	}

	for l1 != nil {
		dummyHead.Next = l1
		dummyHead = dummyHead.Next
		l1 = l1.Next
	}

	for l2 != nil {
		dummyHead.Next = l2
		dummyHead = dummyHead.Next
		l2 = l2.Next
	}

	head = dummy.Next
}
