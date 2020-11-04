package leetcode

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	quickSort(nums, 0, len(nums))
	return
}

func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start

	for j := start; j < end; j++ {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[end], nums[i] = nums[i], nums[end]
	return i
}
