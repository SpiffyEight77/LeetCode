package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	r := strings.Split(s, "")
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			continue
		}
		if s[i] < 97 || s[i] > 122 {
			r[i] = " "
		}
	}

	s = strings.Join(r, " ")
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)

	for i := 0; i < len(s)/2; i++ {
		if string(s[i]) != string(s[len(s)-i-1]) {
			return false
		}
	}
	return true
}

func main() {
	//s := "A man, a plan, a canal: Panama"
	s := "race a car"
	fmt.Println(isPalindrome(s))
}
