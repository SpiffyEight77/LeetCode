package leetcode

import "math"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) <= float64(t) && math.Abs(float64(i-j)) <= float64(k) {
				return true
			}
		}
	}
	return false
}
