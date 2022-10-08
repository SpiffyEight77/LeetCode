package leetcode

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
  isBalanced, _ := maxDepth(root)
  return isBalanced
}

func maxDepth(root *TreeNode) (bool, int) {
  if root == nil {
    return true, 0
  }
  isLeftNodeBalanced, left := maxDepth(root.Left)
  isRightNodeBalanced, right := maxDepth(root.Right)
  if isLeftNodeBalanced && isRightNodeBalanced && abs(left - right) <= 1 {
    return true, max(left, right) + 1
  }
  return false, max(left, right) + 1
}

func abs(x int) int {
  if x > 0 {
    return x
  }
  return -1 * x
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}
