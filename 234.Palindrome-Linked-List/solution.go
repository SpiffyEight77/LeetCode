package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	res := make([]int, 0)
	l := 0
	for head != nil {
		l++
		res = append(res, head.Val)
		head = head.Next
	}

	for i := 0; i <= l/2; i++ {
		if res[i] != res[l-i-1] {
			return false
		}
	}
	return true
}
