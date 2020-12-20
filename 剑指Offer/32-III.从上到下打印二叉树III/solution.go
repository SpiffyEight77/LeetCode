package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	var flag bool

	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			list = append(list, node.Val)
		}
		if flag {
			for i := 0; i < len(list)/2; i++ {
				list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
			}
		}
		flag = !flag
		res = append(res, list)
	}
	return res
}
