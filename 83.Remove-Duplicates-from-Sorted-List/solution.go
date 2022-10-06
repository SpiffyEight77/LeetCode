package leetcode

type ListNode struct {
  Val int
  Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
  if head == nil {
    return nil
  }
  current := head
  for current != nil {
    for current.Next != nil && current.Val == current.Next.Val {
      current.Next = current.Next.Next
    }
    current = current.Next
  }
  return head
}
