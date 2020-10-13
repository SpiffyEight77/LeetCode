package leetcode

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	res := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				res = append(res, strconv.Itoa(node.Val))
				queue = append(queue, node.Left, node.Right)
			} else {
				res = append(res, "#")
			}

		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	res := strings.Split(data, ",")
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(res[0])
	res = res[1:]
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		if res[0] != "#" {
			leftVal, _ := strconv.Atoi(res[0])
			queue[0].Left.Val = leftVal
			queue = append(queue, queue[0].Left)
		}
		if res[1] != "#" {
			rightVal, _ := strconv.Atoi(res[1])
			queue[1].Right.Val = rightVal
			queue = append(queue, queue[0].Right)
		}
		res = res[2:]
		queue = queue[1:]
	}
	return root
}
