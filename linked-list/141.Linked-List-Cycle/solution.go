package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func hasCycleV2(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := map[*ListNode]int{}
	for head != nil && head.Next != nil {
		if m[head] > 0 {
			return true
		}
		m[head]++
		head = head.Next
	}
	return false
}
