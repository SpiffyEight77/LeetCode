package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	res := make([]int, 0)
	res = append(res, root.Val)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp := res[0]
			res = res[1:]
			if node.Left == nil && node.Right == nil && tmp == sum {
				return true
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
				res = append(res, node.Left.Val+tmp)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				res = append(res, node.Right.Val+tmp)
			}
		}
	}
	return false
}

func hasPathSumV2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSumV2(root.Left, sum-root.Val) || hasPathSumV2(root.Right, sum-root.Val)
}
