package leetcode

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
  result := make([][]int, 0)
  if root == nil {
    return result
  }
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  level := 0
  for len(queue) > 0 {
    level++
    length := len(queue)
    list := make([]int, 0)
    for i := 0; i < length; i++ {
      node := queue[0]
      queue = queue[1:]
      if node.Left != nil {
        queue = append(queue, node.Left)
      }
      if node.Right != nil {
        queue = append(queue, node.Right)
      }
      list = append(list, node.Val)
    }
    if level % 2 == 0 {
      for i := 0; i < len(list) / 2; i++ {
        list[i], list[len(list) - i - 1] = list[len(list) - i - 1], list[i]
      }
    }
    result = append(result, list)
  }
  return result
}
