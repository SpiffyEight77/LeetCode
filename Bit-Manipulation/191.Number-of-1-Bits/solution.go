package leetcode

func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		num &= (num - 1)
		res++
	}
	return res
}
