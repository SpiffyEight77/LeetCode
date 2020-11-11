package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0)
	l := len(nums)
	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		right := l - 1
		target := -1 * nums[i]

		for left := i + 1; left < l; left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}

			for left < right && nums[left]+nums[right] > target {
				right--
			}

			if left == right {
				break
			}
			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
			}
		}
	}
	return res
}
