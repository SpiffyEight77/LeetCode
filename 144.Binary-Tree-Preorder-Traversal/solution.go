package leetcode

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
  if root == nil {
    return nil
  }
  result := make([]int, 0)
  preorder(root, &result)
  return result
}

func preorder(root *TreeNode, result *[]int) {
  if root == nil {
    return
  }
  *result = append(*result, root.Val)
  preorder(root.Left, result)
  preorder(root.Right, result)
}
