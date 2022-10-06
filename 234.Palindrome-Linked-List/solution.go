package leetcode

type ListNode struct {
  Val int
  Next *ListNode
}

func isPalindrome(head *ListNode) bool {
  if head == nil {
    return true
  }
  fast, slow := head, head
  for fast != nil && fast.Next != nil {
    fast = fast.Next.Next
    slow = slow.Next
  }
  var prev *ListNode
  for slow != nil {
    temp := slow.Next
    slow.Next = prev
    prev = slow
    slow = temp
  }
  left, right := head, prev
  for right != nil {
    if left.Val != right.Val {
      return false
    }
    left = left.Next
    right = right.Next
  }
  return true
}
