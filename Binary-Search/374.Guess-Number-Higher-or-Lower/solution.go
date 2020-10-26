package leetcode

func guess(num int) int

func guessNumber(n int) int {
	start, end := 0, n
	for start+1 < end {
		mid := start + (end-start)/2
		if guess(mid) == 0 {
			return mid
		}
		if guess(mid) == 1 {
			start = mid
		} else {
			end = mid
		}
	}
	if guess(start) == 0 {
		return start
	}
	return end
}
