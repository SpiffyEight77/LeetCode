package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	pivot := postorder[len(postorder)-1]
	pivotIdx := 0
	for inorder[pivotIdx] != pivot {
		pivotIdx++
	}

	left := buildTree(inorder[:pivotIdx], postorder[:pivotIdx])
	right := buildTree(inorder[pivotIdx+1:], postorder[pivotIdx:len(postorder)-1])

	return &TreeNode{
		Val:   pivot,
		Left:  left,
		Right: right,
	}
}
