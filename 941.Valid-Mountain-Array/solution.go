package leetcode

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	start, end := 0, len(A)-1
	for start < end && A[start] < A[start+1] {
		start++
	}
	for start < end && A[end] < A[end-1] {
		end--
	}
	if start == 0 || end == len(A)-1 {
		return false
	}
	return true
}
