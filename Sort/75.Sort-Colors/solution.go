package leetcode

import "sort"

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return
}
