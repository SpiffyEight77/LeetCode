package leetcode

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
  prices := []int{7, 1, 5, 3, 6, 4}
  result := 5
  if ans := MaxProfit(prices); ans != result {
    t.Errorf("Wrong answer")
  }
}
