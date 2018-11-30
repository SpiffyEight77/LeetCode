package main

import "fmt"

// v2
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	res := 0
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]]++
	}
	for key,_ := range dict {
		v := dict[key+k]
		if k == 0 && v == 1 || v == 0 {
			continue
		}
		res++
	}
	return res
}

func main() {
	nums := []int{3, 1, 4, 1, 5}
	k := 2
	fmt.Println(findPairs(nums, k))
}
