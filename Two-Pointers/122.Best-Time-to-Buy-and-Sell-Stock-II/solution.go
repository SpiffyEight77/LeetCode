package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
