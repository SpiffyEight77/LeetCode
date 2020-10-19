package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m > n {
		return head
	}
	dummy := &ListNode{0, head}
	head = dummy

	var prev *ListNode
	var i = 0
	for i < m {
		prev = head
		head = head.Next
		i++
	}

	var j = i
	var next *ListNode
	var mid = head
	for head != nil && j <= n {
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		j++
	}

	prev.Next = next
	mid.Next = head

	return dummy.Next
}
