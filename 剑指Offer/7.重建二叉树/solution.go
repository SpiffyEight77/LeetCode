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
	index, rootVal := 0, preorder[0]

	for index = 0; index < len(inorder); index++ {
		if inorder[index] == rootVal {
			break
		}
	}

	root := &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
	return root
}
