package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1

	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[start] < nums[mid] {
			if target <= nums[mid] && nums[start] <= target {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
