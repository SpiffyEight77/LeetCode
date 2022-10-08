package leetcode

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  stopped := false
  for len(queue) > 0 {
    length := len(queue)
    for i := 0; i < length; i++ {
      node := queue[0]
      queue = queue[1:]
      if node == nil {
        stopped = true
      }
      if node != nil && stopped {
        return false
      }
      if node != nil {
        queue = append(queue, node.Left)
        queue = append(queue, node.Right)
      }
    }
  }
  return true
}
