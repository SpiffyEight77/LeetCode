package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftDepth := dfs(root.Left, res)
	rightDepth := dfs(root.Right, res)
	*res = max(leftDepth+rightDepth, *res)
	return max(leftDepth, rightDepth) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
