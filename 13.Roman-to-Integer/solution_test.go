package leetcode

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
  s := "MCMXCIV"
  result := 1994
  if ans := RomanToInt(s); ans != result {
    t.Errorf("Wrong answer")
  }
}
