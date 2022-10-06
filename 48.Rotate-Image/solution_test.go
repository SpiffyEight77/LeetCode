package leetcode

import (
	"testing"
)

func TestRotate(t *testing.T) {
  matrix := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
  matrix2 := [][]int{{1,2,3},{4,5,6},{7,8,9}}
  // result := [][]int{{15,13,2,5},{14,3,4,1},{12,6,8,9},{16,7,10,11}}
  Rotate(matrix)
  Rotate(matrix2)
}
