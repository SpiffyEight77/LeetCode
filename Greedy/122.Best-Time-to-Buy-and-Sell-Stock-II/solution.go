package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var (
		res    = 0
		index  = 0
		target = 0
	)

	for index < len(prices)-1 {
		for index < len(prices)-1 && prices[index] >= prices[index+1] {
			index++
		}
		target = prices[index]
		for index < len(prices)-1 && prices[index] <= prices[index+1] {
			index++
		}
		res += prices[index] - target
	}
	return res
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
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
