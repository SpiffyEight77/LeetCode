package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	res := make([]int, 0)
	inOrder(root, &res)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

func isValidBSTV2(root *TreeNode) bool {
	if root == nil {
		return false
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
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}
