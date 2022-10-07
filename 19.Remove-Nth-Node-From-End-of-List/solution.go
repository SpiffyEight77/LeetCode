package leetcode

type ListNode struct {
  Val int
  Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  dummyHead := &ListNode{
    Val: -1,
    Next: head,
  }
  left := dummyHead
  right := head
  for n > 0 && right != nil {
    n--
    right = right.Next
  }
  for right != nil {
    left = left.Next
    right = right.Next
  }
  left.Next = left.Next.Next
  return dummyHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
  dummyHead := &ListNode{
    Val: -1,
    Next: head,
  }
  cnt := 0
  ptr := head
  for ptr != nil {
    ptr = ptr.Next
    cnt++
  }
  cnt -= n
  ptr = dummyHead
  for cnt > 0 {
    cnt--
    ptr = ptr.Next
  }
  ptr.Next = ptr.Next.Next
  return dummyHead.Next
}


