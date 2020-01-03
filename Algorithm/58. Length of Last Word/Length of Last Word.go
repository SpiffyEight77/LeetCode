package main

import (
	"strings"
)

//v1 O(N) O(N)
func lengthOfLastWord(s string) int {
	res := strings.Split(s, " ")
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] != ""{
			return len(res[i])
		}
	}
	return 0
}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	res := strings.Split(s, " ")
	return len(res[len(res)-1])
}
