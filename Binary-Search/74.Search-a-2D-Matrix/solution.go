package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix)
	col := len(matrix[0])
	start, end := 0, row*col-1
	for start+1 < end {
		mid := start + (end-start)/2
		if matrix[mid/col][mid%col] == target {
			return true
		}
		if matrix[mid/col][mid%col] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target {
		return true
	}

	return false
}
