package leetcode

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	start, end := 0, num/2
	for start+1 < end {
		mid := start + (end-start)/2
		if mid*mid == num {
			return true
		}
		if mid*mid < num {
			start = mid
		} else {
			end = mid
		}
	}
	if start*start == num || end*end == num {
		return true
	}
	return false
}
