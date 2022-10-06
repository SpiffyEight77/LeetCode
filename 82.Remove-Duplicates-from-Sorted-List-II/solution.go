package leetcode

type ListNode struct {
  Val int
  Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
   if head == nil || head.Next == nil {
     return head
   }
   dummyHead := &ListNode{
     Val: -1,
     Next: head,
   }
   prev := dummyHead
   for prev.Next != nil && prev.Next.Next != nil {
     if prev.Next.Val == prev.Next.Next.Val {
       rmVal := prev.Next.Val
       for prev.Next != nil && prev.Next.Val == rmVal {
         prev.Next = prev.Next.Next
       }
     } else {
       prev = prev.Next
     }
   }
   return dummyHead.Next
}

func deleteDuplicatesThreePointers(head *ListNode) *ListNode {
   if head == nil || head.Next == nil {
     return head
   }
   dummyHead := &ListNode{
     Val: -1,
     Next: head,
   }
   prev := dummyHead
   for head != nil {
     if head.Next != nil && head.Next.Val == head.Val {
       for head.Next != nil && head.Next.Val == head.Val {
         head = head.Next
       }
       prev.Next = head.Next
     } else {
       prev = prev.Next
     }
     head = head.Next
   }
   return dummyHead.Next
}

