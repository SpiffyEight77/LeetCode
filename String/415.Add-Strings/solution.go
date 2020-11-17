package leetcode

import "strconv"

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	res := ""

	for i >= 0 || j >= 0 || carry != 0 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		carry += x + y
		res = strconv.Itoa(carry%10) + res
		carry /= 10

		i--
		j--
	}
	return res
}
