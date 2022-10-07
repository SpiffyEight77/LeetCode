package leetcode

type ListNode struct {
  Val int
  Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {
    return nil
  }
  fast, slow := head, head
  for fast.Next != nil && fast.Next.Next != nil {
    fast = fast.Next.Next
    slow = slow.Next
    if fast == slow {
      break
    }
  }
  slow = head
  for slow != fast {
    slow = slow.Next
    fast = fast.Next
  }
  return slow
}
