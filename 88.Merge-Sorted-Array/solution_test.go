package leetcode

import (
	"testing"
)

func TestMerge(t *testing.T) {
  nums1 := []int{1, 2, 3, 0, 0, 0}
  nums2 := []int{2, 5, 6}
  m, n := 3, 3
  // result := []int{1, 2, 2, 3, 5, 6}
  Merge(nums1 ,m , nums2, n)
}
