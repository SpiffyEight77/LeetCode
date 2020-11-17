package leetcode

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	f := make([]int, len(nums))
	f[0] = 0

	for i := 1; i < len(nums); i++ {
		idx := 0
		for idx < len(nums) && idx+nums[idx] < i {
			idx++
		}
		f[i] = f[idx] + 1
	}

	return f[len(nums)-1]
}

func jumpV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var rightMost int
	var res int
	var end int

	for i := 0; i < len(nums)-1; i++ {
		rightMost = max(nums[i]+i, rightMost)
		if i == end {
			end = rightMost
			res++
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
