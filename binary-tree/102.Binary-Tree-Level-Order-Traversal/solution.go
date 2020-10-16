package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	levelOrderDFS(root, 0)
	return res
}

func levelOrderDFS(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	levelOrderDFS(root.Left, level+1)
	levelOrderDFS(root.Right, level+1)

	return
}

func levelOrderV2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, list)
	}
	return res
}
