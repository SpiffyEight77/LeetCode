package leetcode

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	f := make([]bool, len(nums))
	f[0] = true

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if f[j] == true && nums[j]+j >= i {
				f[i] = true
			}
		}
	}
	return f[len(nums)-1]
}

func canJumpV2(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	var rightMost int

	for i := 0; i < len(nums); i++ {
		if i > rightMost {
			return false
		} else {
			rightMost = max(rightMost, i+nums[i])
			if rightMost >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
