package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	dummy *ListNode
	size  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{dummy: &ListNode{Val: 0}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	head := this.dummy.Next
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	prev := this.dummy
	for prev.Next != nil {
		prev = prev.Next
	}
	prev.Next = &ListNode{Val: val, Next: prev.Next}
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	prev := this.dummy
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = &ListNode{val, prev.Next}
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := this.dummy
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	this.size--
}

/**
* Your MyLinkedList object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Get(index);
* obj.AddAtHead(val);
* obj.AddAtTail(val);
* obj.AddAtIndex(index,val);
* obj.DeleteAtIndex(index);
 */
