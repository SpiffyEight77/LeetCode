package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postorderTraversalDFS(root, &res)
	return res
}

func postorderTraversalDFS(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postorderTraversalDFS(root.Left, res)
	postorderTraversalDFS(root.Right, res)
	*res = append(*res, root.Val)
}

func postorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return res
}
