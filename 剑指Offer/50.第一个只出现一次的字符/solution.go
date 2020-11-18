package leetcode

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	freq := [256]int{}

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if freq[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
