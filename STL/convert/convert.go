package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "12345"
	num := int(s[0] - '0')
	str := string(s[0])
	b := byte(num + '0')

	n, _ := strconv.Atoi(s)
	ss := strconv.Itoa(n)
	fmt.Printf("%d %s %c %d %s\n", num, str, b, n, ss)
}
