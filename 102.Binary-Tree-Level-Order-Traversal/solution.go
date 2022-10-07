package leetcode

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
  result := make([][]int, 0)
  if root == nil {
    return result
  } 
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  for len(queue) > 0 {
    list := make([]int, 0)
    length := len(queue)
    for i := 0; i < length; i++ {
      node := queue[0]
      queue = queue[1:]
      list = append(list, node.Val)
      if node.Left != nil {
        queue = append(queue, node.Left)
      }
      if node.Right != nil {
        queue = append(queue, node.Right)
      }
    }
    result = append(result, list)
  }
  return result
}
