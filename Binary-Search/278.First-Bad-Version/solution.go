package leetcode

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	start, end := 0, n
	for start+1 < end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid
		}
	}
	if isBadVersion(start) {
		return start
	}
	return end
}
