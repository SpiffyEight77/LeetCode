package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	queue = append(queue, root)

	for len(queue) > 0 {
		p := queue[0]
		q := queue[1]
		queue = queue[2:]

		if p == nil && q == nil {
			continue
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		queue = append(queue, p.Left)
		queue = append(queue, q.Right)
		queue = append(queue, p.Right)
		queue = append(queue, q.Left)

	}
	return true
}
