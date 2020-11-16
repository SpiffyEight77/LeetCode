package leetcode

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	freq := make(map[int]int, 0)
	res := make([]int, 0)

	for _, v := range arr1 {
		freq[v]++
	}

	for _, v := range arr2 {
		for freq[v] != 0 {
			res = append(res, v)
			freq[v]--
		}
	}

	tmp := make([]int, 0)
	for i := 0; i < len(arr1); i++ {
		for freq[arr1[i]] != 0 {
			tmp = append(tmp, arr1[i])
			freq[arr1[i]]--
		}
	}

	sort.Ints(tmp)
	res = append(res, tmp...)

	return res
}

func relativeSortArrayV2(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return arr1
	}

	freq := [1001]int{}
	res := make([]int, 0)

	for _, v := range arr1 {
		freq[v]++
	}

	for _, v := range arr2 {
		for ; freq[v] != 0; freq[v]-- {
			res = append(res, v)
		}
	}

	for k, v := range freq {
		for ; v != 0; v-- {
			res = append(res, k)
		}
	}
	return res
}
