package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	toggle := false

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
		if toggle {
			reverse(list)
		}
		res = append(res, list)
		toggle = !toggle
	}
	return res
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
