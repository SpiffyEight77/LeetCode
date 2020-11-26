package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	head = dfs(head)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func dfs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := dfs(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}
