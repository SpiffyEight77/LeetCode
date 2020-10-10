package main

import (
	"fmt"
	"sort"
)

func main() {
	sort.Ints([]int{})

	sort.Strings([]string{})

	s := []int{5, 7, 1, 2, 3, 4}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	fmt.Println(s)
}
