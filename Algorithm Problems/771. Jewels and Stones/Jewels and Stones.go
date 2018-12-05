package main

import "fmt"

// v1
func numJewelsInStones(J string, S string) int {
	dict := make(map[string]bool)
	var sum int

	for i := 0; i < len(J); i++ {
		dict[string(J[i])] = true
	}

	for i := 0; i < len(S); i++ {
		if _, ok := dict[string(S[i])]; ok {
			sum++
		}
	}

	return sum

}

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}
