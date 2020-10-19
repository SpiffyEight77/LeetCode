package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	nums := make([]int, 0)
	for head.Next != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = dfs(nums, left, mid-1)
	root.Right = dfs(nums, mid+1, right)
	return root
}
