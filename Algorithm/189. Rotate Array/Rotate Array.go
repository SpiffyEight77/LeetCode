package main

//v1
func rotate(nums []int, k int) {
	l := len(nums)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[(i+k)%l] = nums[i]
	}
	for i := 0; i < l; i++ {
		nums[i] = res[i]
	}
}

