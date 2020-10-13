package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	res  []int
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	res := make([]int, 0)
	inOrder(root, &res)
	return BSTIterator{res: res, root: root}
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	val := this.res[0]
	this.res = this.res[1:]
	return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.res) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
