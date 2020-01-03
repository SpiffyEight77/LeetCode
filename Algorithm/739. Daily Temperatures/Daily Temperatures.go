package main

import "fmt"

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	index := make([]int, 0, len(T))

	for i := 0; i < len(T); i++ {
		if len(index) != 0 && T[i] > T[index[len(index)-1]] {
			for {
				if len(index) == 0 || T[i] <= T[index[len(index)-1]] {
					break
				}
				res[index[len(index)-1]] = i - index[len(index)-1]
				index = append(index[:len(index)-1])
			}
		}
		index = append(index, i)
	}
	return res
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}
