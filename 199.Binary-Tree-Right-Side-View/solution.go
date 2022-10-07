package leetcode

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
  result := make([]int, 0)
  if root == nil {
    return result
  }
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  for len(queue) > 0 {
    var rightSideNode *TreeNode
    length := len(queue)
    for i := 0; i < length; i++ {
      node := queue[0]
      queue = queue[1:]
      if node != nil {
        rightSideNode = node
        queue = append(queue, node.Left)
        queue = append(queue, node.Right)
      }
    }
    if rightSideNode != nil {
      result = append(result, rightSideNode.Val)
    }
  }
  return result
}

