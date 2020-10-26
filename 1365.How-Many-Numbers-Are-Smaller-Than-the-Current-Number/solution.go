package leetcode

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}
