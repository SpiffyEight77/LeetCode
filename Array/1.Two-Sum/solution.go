package leetcode

import "sort"

func twoSum(nums []int, target int) []int {
	freq := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if v, ok := freq[target-nums[i]]; ok {
			return []int{v, i}
		} else {
			freq[nums[i]] = i
		}
	}
	return nil
}

type SortedNum struct {
	Val   int
	Index int
}

func twoSumV2(nums []int, target int) []int {
	sortedNum := make([]SortedNum, len(nums))
	for i := 0; i < len(sortedNum); i++ {
		sortedNum[i].Val = nums[i]
		sortedNum[i].Index = i
	}
	sort.Slice(sortedNum, func(i, j int) bool {
		return sortedNum[i].Val < sortedNum[j].Val
	})

	left, right := 0, len(sortedNum)-1
	for left < right {
		if sortedNum[left].Val+sortedNum[right].Val == target {
			if sortedNum[left].Index > sortedNum[right].Index {
				return []int{sortedNum[right].Index, sortedNum[left].Index}
			}
			return []int{sortedNum[left].Index, sortedNum[right].Index}
		} else if sortedNum[left].Val+sortedNum[right].Val > target {
			right--
		} else {
			left++
		}
	}
	return nil
}
