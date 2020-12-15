package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][1] {
			continue
		} else if res[len(res)-1][1] >= intervals[i][0] && res[len(res)-1][1] < intervals[i][1] {
			res[len(res)-1][1] = intervals[i][1]
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
