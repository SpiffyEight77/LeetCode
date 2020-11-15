package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	node := &TreeNode{}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node = queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, node.Val)
	}
	return res
}

func rightSideViewV2(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, root.Val)
	}
	dfs(root.Right, level+1, res)
	dfs(root.Left, level+1, res)
}
