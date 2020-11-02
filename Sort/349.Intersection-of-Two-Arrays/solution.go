package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	count := map[int]int{}
	count1 := map[int]int{}
	count2 := map[int]int{}
	nums11 := []int{}
	nums22 := []int{}
	res := make([]int, 0)

	for _, v := range nums1 {
		count1[v]++
	}

	for _, v := range nums2 {
		count2[v]++
	}

	for i := range count1 {
		nums11 = append(nums11, i)
	}

	for i := range count2 {
		nums22 = append(nums22, i)
	}

	nums11 = append(nums11, nums22...)

	for _, v := range nums11 {
		count[v]++
	}

	for i, v := range count {
		if v > 1 {
			res = append(res, i)
		}
	}

	return res
}
