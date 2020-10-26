package leetcode

func mySqrt(x int) int {
	start, end := 0, x
	for start+1 < end {
		mid := start + (end-start)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			start = mid
		} else {
			end = mid
		}
	}
	if end*end <= x {
		return end
	}
	return start
}
