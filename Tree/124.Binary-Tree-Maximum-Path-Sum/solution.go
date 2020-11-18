package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftVal := max(dfs(root.Left, res), 0)
	rightVal := max(dfs(root.Right, res), 0)

	*res = max(*res, root.Val+leftVal+rightVal)

	return root.Val + max(leftVal, rightVal)

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
