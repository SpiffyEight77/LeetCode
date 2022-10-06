package leetcode

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
  // height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
  height := []int{1, 1}
  result := 1
  if ans := MaxArea(height); ans != result {
    t.Errorf("Wrong answer")
  }
}
