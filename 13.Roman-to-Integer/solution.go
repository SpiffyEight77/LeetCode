package leetcode

func RomanToInt(s string) int {
  sum := 0
  romanToIntMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100,
  'D': 500, 'M': 1000}
  for index, value := range s {
    sum += romanToIntMap[value]
    if index > 0 && s[index - 1] == 'I' && (value == 'V' || value == 'X') {
      sum -= 1 * 2
    }
    if index > 0 && s[index - 1] == 'X' && (value == 'L' || value == 'C') {
      sum -= 10 * 2
    }
    if index > 0 && s[index - 1] == 'C' && (value == 'D' || value == 'M') {
      sum -= 100 * 2
    }
  }
  return sum
}
