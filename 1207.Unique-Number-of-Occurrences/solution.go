package leetcode

func uniqueOccurrences(arr []int) bool {
	nums := map[int]int{}
	count := map[int]int{}

	for _, v := range arr {
		nums[v]++
	}

	for _, v := range nums {
		count[v]++
	}

	for _, v := range count {
		if v > 1 {
			return false
		}
	}
	return true
}
