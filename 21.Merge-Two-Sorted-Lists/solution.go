package leetcode

type ListNode struct {
  Val int
  Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
      return list2
    }
    if list2 == nil {
      return list1
    }
    if list1 == nil && list2 == nil {
      return nil
    }
    dummyHead := &ListNode{
      Val: -1,
    }
    head := dummyHead
    for list1 != nil && list2 != nil {
      if list1.Val > list2.Val {
        head.Next = list2
        list2 = list2.Next
      } else {
        head.Next = list1
        list1 = list1.Next
      }
      head = head.Next
    }
    if list1 == nil {
      head.Next = list2
    }
    if list2 == nil {
      head.Next = list1
    }
    return dummyHead.Next
}
