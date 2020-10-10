package main

import "fmt"

func main() {
	n := 10
	x := 100
	a := make([]int, n)
	a[n-1] = x
	fmt.Println(a)

	i := 3
	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]
	fmt.Println(a)

	b := make([]int, 0)
	b = append(b, x)
	fmt.Println(b)
}
