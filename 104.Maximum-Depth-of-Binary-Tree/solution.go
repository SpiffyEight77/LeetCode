package leetcode

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}

func maxDepth(root *TreeNode) int {
  level := 0
  if root == nil {
    return level
  }
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  for len(queue) > 0 {
    length := len(queue)
    for i := 0; i < length; i++ {
      node := queue[0]
      queue = queue[1:]
      if node.Left != nil {
        queue = append(queue, node.Left)
      }
      if node.Right != nil {
        queue = append(queue, node.Right)
      }
    }
    level++
  }
  return level
}

func maxDepthDFS(root *TreeNode) int {
  if root == nil {
    return 0
  }
  return max(maxDepthDFS(root.Left), maxDepthDFS(root.Right)) + 1
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}
