package leetcode

type ListNode struct {
  Val int
  Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
  if headA == nil || headB == nil {
    return nil
  }
  pointerA, pointerB := headA, headB
  for pointerA != pointerB {
    if pointerA == nil {
      pointerA = headB
    } else {
      pointerA = pointerA.Next
    }
    if pointerB == nil {
      pointerB = headA
    } else {
      pointerB = pointerB.Next
    }
  }
  return pointerA
}


