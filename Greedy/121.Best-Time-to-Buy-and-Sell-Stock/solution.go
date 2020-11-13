package leetcode

import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var res int
	l, r := 0, 1
	for r < len(prices) {
		if prices[r] < prices[l] {
			l = r
		}
		if prices[r] > prices[l] {
			res = max(res, prices[r]-prices[l])
		}
		r++
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfitV2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
