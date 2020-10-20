package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTreeV2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, p)
	queue = append(queue, q)

	for len(queue) > 0 {
		p = queue[0]
		q = queue[1]
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
		queue = append(queue, q.Left)
		queue = append(queue, p.Right)
		queue = append(queue, q.Right)
	}
	return true
}
