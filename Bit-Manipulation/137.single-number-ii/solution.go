package leetcode

func singleNumber(nums []int) int {
	var res int
	for i := 0; i < 64; i++ {
		var sum int
		for j := 0; j < len(nums); j++ {
			sum += nums[j] >> i & 1
		}
		res ^= sum % 3 << i
	}
	return res
}
