package leetcode

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	got := QuickSort([]int{7, 1, 2, 6, 4, 5, 3})
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
