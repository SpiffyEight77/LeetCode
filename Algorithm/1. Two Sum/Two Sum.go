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

//v2 hash table O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			slice := []int{v, i}
			return slice
		}
		m[target-nums[i]] = i
	}
	return nil
}

//v3 暴力
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res := []int{i, j}
				return res
			}
		}
	}
	return nil
}
