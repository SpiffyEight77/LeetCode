package leetcode

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return nil
	}
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	res = append(res, intervals[0])
	idx := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[idx][1] {
			res = append(res, intervals[i])
			idx++
		} else {
			res[idx][1] = max(res[idx][1], intervals[i][1])
		}
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
