package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		res = max(nums[i], res)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
