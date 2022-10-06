package leetcode

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
  matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
  target := 20
  result := false
  if ans := SearchMatrix(matrix, target); ans != result {
    t.Errorf("Wrong answer")
  }
}
