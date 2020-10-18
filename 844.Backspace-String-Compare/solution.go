package leetcode

func backspaceCompare(S string, T string) bool {
	S = check(S)
	T = check(T)
	return S == T
}

func check(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[i] == '#' {
			stack = stack[:len(stack)-1]
			continue
		} else if s[i] != '#' {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
