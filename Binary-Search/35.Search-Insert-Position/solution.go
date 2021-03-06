package leetcode

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			end = mid
		}
		if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] < target {
		return end + 1
	}
	return 0
}
