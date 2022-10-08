package leetcode

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return nil
  }
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  for len(queue) > 0 {
    length := len(queue)
    for i := 0; i < length; i++ {
      node := queue[0]
      queue = queue[1:]
      node.Left, node.Right = node.Right, node.Left 
      if node.Left != nil {
        queue = append(queue, node.Left)
      }
      if node.Right != nil {
        queue = append(queue, node.Right)
      }
    }
  }
  return root
}
