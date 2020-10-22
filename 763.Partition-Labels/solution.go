package leetcode

func partitionLabels(S string) []int {
	lastPos := [26]int{}
	for i, c := range S {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	partition := make([]int, 0)
	for i, c := range S {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return partition
}
