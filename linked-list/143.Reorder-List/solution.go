package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	var prev *ListNode

	for mid != nil {
		temp := mid.Next
		mid.Next = prev
		prev = mid
		mid = temp
	}
	slow.Next = nil

	l1 := head
	l2 := prev
	dummy := &ListNode{0, head}
	head = dummy
	toggle := true

	for l1 != nil && l2 != nil {
		if toggle {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
		toggle = !toggle
	}

	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}

	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}

	head = dummy.Next
}
