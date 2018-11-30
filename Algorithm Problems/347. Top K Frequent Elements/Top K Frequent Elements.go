package main

import (
	"fmt"
)

// v1
func topKFrequent(nums []int, k int) []int {
	l := len(nums)
	res := make([]int, 0, k)
	dict := make(map[int]int, l)
	array := make([]int, 0, l)
	for _, v := range nums {
		dict[v]++
	}

	for k, _ := range dict {
		array = append(array, k)
	}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if dict[array[i]] < dict[array[j]] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	//fmt.Println(array)
	//fmt.Println(dict)
	for i := 0; i < k; i++ {
		res = append(res, array[i])
	}
	return res
}

func main() {
	//nums := []int{1, 1, 1, 2, 2, 3}
	nums := []int{1}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}
