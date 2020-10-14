package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ResultType struct {
	height int
	valid  bool
}

func isBalanced(root *TreeNode) bool {
	return maxHeight(root).valid
}

func maxHeight(root *TreeNode) (res ResultType) {
	if root == nil {
		res.valid = true
		res.height = 0
		return
	}

	left := maxHeight(root.Left)
	right := maxHeight(root.Right)

	if left.valid && right.valid && abs(left.height, right.height) <= 1 {
		res.valid = true
	}
	res.height = Max(left.height, res.height)
	return
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
