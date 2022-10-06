package leetcode

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
  matrix := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
  result := []int{1,2,3,6,9,8,7,4,5}
  if ans := SpiralOrder(matrix); !reflect.DeepEqual(ans, result) {
    t.Errorf("Wrong answer")
  }
}
