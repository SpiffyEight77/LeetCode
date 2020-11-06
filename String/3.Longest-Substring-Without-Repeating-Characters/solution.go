package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := map[byte]int{}
	right, ans, l := -1, 0, len(s)
	for i := 0; i < l; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < l && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		ans = max(ans, right-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
