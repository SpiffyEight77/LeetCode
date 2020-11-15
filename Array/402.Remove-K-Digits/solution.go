package leetcode

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	stack := make([]byte, 0)

	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]

	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:]
	}
	return string(stack)
}
