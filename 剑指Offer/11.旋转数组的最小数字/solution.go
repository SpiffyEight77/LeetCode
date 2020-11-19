package leetcode

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	start, end := 0, len(numbers)-1

	for start+1 < end {
		for start < end && numbers[start] == numbers[start+1] {
			start++
		}

		for start < end && numbers[end] == numbers[end-1] {
			end--
		}

		mid := start + (end-start)/2
		if numbers[mid] >= numbers[end] {
			start = mid
		} else {
			end = mid
		}
	}
	if numbers[start] > numbers[end] {
		return numbers[end]
	}
	return numbers[start]
}
