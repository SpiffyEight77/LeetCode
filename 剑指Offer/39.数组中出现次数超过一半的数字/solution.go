package leetcode

func majorityElement(nums []int) int {
	count := make(map[int]int, len(nums))
	for _, v := range nums {
		count[v]++
		if count[v] > len(nums)/2 {
			return v
		}
	}
	return -1
}
