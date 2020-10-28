package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)

	for i := 0; i < row; i++ {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}
