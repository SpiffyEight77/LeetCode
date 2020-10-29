package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	res := make([]int, 0)
	dfs(root, 0, &res)
	var sum int
	for _, v := range res {
		sum += v
	}
	return sum
}

func dfs(root *TreeNode, num int, res *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*res = append(*res, num*10+root.Val)
	}

	dfs(root.Left, num*10+root.Val, res)
	dfs(root.Right, num*10+root.Val, res)
}
