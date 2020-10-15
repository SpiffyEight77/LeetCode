package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(root, &res)
	return res
}

func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
}

func preorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return res
}

func preorderTraversalV3(root *TreeNode) []int {
	return divideAndConquer(root)
}

func divideAndConquer(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)

	return res
}
