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

func maxSubArrayV2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	f := make([]int, len(nums))
	f[0] = nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		f[i] = max(nums[i], f[i-1]+nums[i])
		res = max(res, f[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
