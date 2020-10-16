package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorderTraversalDFS(root, &res)
	return res
}

func inorderTraversalDFS(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorderTraversalDFS(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalDFS(root.Right, res)
}

func inorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}

	return res
}
