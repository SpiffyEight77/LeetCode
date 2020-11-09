package leetcode

import "sort"

func kClosest(points [][]int, K int) [][]int {
	if len(points) == 0 {
		return nil
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:K]
}
