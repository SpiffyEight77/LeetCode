package leetcode

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}

		mid := start + (end-start)/2
		if nums[end] >= nums[mid] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}
