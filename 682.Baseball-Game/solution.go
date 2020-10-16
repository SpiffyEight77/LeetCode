package leetcode

import "strconv"

func calPoints(ops []string) int {
	stack := make([]int, 0)
	for i := range ops {
		switch ops[i] {
		default:
			num, _ := strconv.Atoi(ops[i])
			stack = append(stack, num)
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		}
	}

	res := 0
	for i := range stack {
		res += stack[i]
	}
	return res
}
