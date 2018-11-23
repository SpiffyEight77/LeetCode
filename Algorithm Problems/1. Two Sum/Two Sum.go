package main

//v1 two points O(n*n)
func twoSum(nums []int, target int) []int {
	i, j := 1, 0
	slice := make([]int, 0, 2)
	for {
		if nums[i]+nums[j] == target {
			slice = append(slice, j)
			slice = append(slice, i)
			return slice
		}
		i++
		if i == len(nums) {
			j++
			i = j + 1
		}
		if j == len(nums) {
			return slice
		}
	}
}

//v2 map O(n)
func twoSum(nums []int, target int) []int {
	slice := make([]int, 2, 2)
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			slice = []int{v, i}
			break
		} else {
			m[target-nums[i]] = i
		}
	}
	return slice
}
