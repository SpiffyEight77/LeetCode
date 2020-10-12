package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	var index int
	for index = 0; index < len(inorder); index++ {
		if inorder[index] == rootVal {
			break
		}
	}

	left := buildTree(preorder[1:index+1], inorder[:index])
	right := buildTree(preorder[index+1:], inorder[index+1:])
	root := &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right,
	}

	return root
}
