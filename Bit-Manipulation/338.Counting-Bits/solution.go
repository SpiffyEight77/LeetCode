package leetcode

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = count1(i)
	}
	return res
}

func count1(x int) int {
	count := 0
	for x != 0 {
		x &= (x - 1)
		count++
	}
	return count
}
