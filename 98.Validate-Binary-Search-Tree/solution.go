package leetcode

import "math"

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
  if root == nil {
    return false
  }
  return helper(root, math.MinInt32 - 1, math.MaxInt32 + 1)
}

func helper(node *TreeNode, leftVal, rightVal int) bool {
  if node == nil {
    return true
  }
  if node.Val >= rightVal || node.Val <= leftVal {
    return false
  }
  return helper(node.Left, leftVal, node.Val) && helper(node.Right, node.Val, rightVal)
}
