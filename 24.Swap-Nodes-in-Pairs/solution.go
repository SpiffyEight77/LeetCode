package leetcode

type ListNode struct {
  Val int
  Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
 if head == nil || head.Next == nil {
   return head
 }
 dummyHead := &ListNode{
   Val: -1,
   Next: head,
 }
 prev, current := head, head.Next
 return head
}
