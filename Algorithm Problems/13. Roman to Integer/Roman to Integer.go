package main

import (
	"fmt"
)

//v1
//func romanToInt(s string) int {
//	tmpMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
//	score := 0
//	for i := 0; i < len(s); i++ {
//		if i != len(s)-1 && string(s[i]) == "I" && string(s[i+1]) == "V" {
//			score += 4
//			i++
//		} else if i != len(s)-1 && string(s[i]) == "I" && string(s[i+1]) == "X" {
//			score += 9
//			i++
//		} else if i != len(s)-1 && string(s[i]) == "X" && string(s[i+1]) == "L" {
//			score += 40
//			i++
//		} else if i != len(s)-1 && string(s[i]) == "X" && string(s[i+1]) == "C" {
//			score += 90
//			i++
//		} else if i != len(s)-1 && string(s[i]) == "C" && string(s[i+1]) == "D" {
//			score += 400
//			i++
//		} else if i != len(s)-1 && string(s[i]) == "C" && string(s[i+1]) == "M" {
//			score += 900
//			i++
//		} else {
//			score += tmpMap[string(s[i])]
//		}
//	}
//	return score
//}

//v2
func romanToInt(s string) int {
	score := 0
	dict := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	res := []rune(s)

	for i := 0; i < len(s)-1; i++ {
		if dict[res[i]] < dict[res[i+1]] {
			score -= dict[res[i]]
		} else {
			score += dict[res[i]]
		}
	}
	score += dict[res[len(s)-1]]
	return score
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
