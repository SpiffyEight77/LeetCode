package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
  nums := []int{2, 7, 11, 15}
  target := 9
  result := []int{0, 1}
  if ans := TwoSum(nums, target); !reflect.DeepEqual(ans, result) {
    t.Errorf("Wrong answer")
  }
}
