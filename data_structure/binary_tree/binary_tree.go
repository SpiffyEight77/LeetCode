package main

import (
	"fmt"
)

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

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		res = append(res, list)
	}
	return res
}

func preorderDivide(root *Node) []int {
	return preorderDivideTraversal(root)
}

func preorderDivideTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	left := preorderDivideTraversal(root.Left)
	right := preorderDivideTraversal(root.Right)

	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

func preorderDFSTraversal(root *Node) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(root *Node, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

func levelOrderDFSTraversal(root *Node) [][]int {
	return levelOrderDFS(root, 0, [][]int{})
}

func levelOrderDFS(root *Node, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) == level {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	res = levelOrderDFS(root.Left, level+1, res)
	res = levelOrderDFS(root.Right, level+1, res)
	return res
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftdepth := maxDepth(root.Left)
	rightdepth := maxDepth(root.Right)
	if leftdepth > rightdepth {
		return leftdepth + 1
	} else {
		return rightdepth + 1
	}
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

	fmt.Println(levelOrder(root))

	fmt.Println(levelOrderDFSTraversal(root))

	fmt.Println("max depth:", maxDepth(root))
}
