package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

func preorderTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Println(root.Val)
	inorder(root.Right)
}

func inorderTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}

	return result
}

func postorder(root *Node) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Println(root.Val)
}

func postorderTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*Node, 0)
	var lastvisit *Node

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastvisit {
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			lastvisit = node
		} else {
			root = node.Right
		}
	}
	return res
}

func main() {

	root := &Node{Val: 3}
	root.Left = &Node{Val: 0}
	root.Left.Left = &Node{Val: 5}
	root.Left.Right = &Node{Val: 1}
	root.Right = &Node{Val: 4}
	root.Right.Left = &Node{Val: 8}
	root.Right.Right = &Node{Val: 10}

	res := preorderTraversal(root)
	fmt.Println(res)

	res = inorderTraversal(root)
	fmt.Println(res)

	postorderTraversal(root)
}
