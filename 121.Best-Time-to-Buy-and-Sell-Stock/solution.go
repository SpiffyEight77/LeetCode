package leetcode

func MaxProfit(prices []int) int {
  profit := 0
  for lowestPrice, index:=prices[0], 1; index < len(prices); index++ {
    profit = max(profit, prices[index] - lowestPrice)
    lowestPrice = min(lowestPrice, prices[index])
  }
  return profit
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func min(x, y int) int {
  if x > y {
    return y
  }
  return x
}

