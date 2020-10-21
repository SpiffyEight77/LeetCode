package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		var nodeX, nodeY int
		var flagX, flagY bool
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
				if node.Left.Val == x {
					nodeX = node.Val
					flagX = true
				}
				if node.Left.Val == y {
					nodeY = node.Val
					flagY = true
				}
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				if node.Right.Val == y {
					nodeY = node.Val
					flagY = true
				}
				if node.Right.Val == x {
					nodeX = node.Val
					flagX = true
				}
			}

		}
		if flagX == true && flagY == true {
			return nodeX != nodeY
		}
	}
	return false
}
