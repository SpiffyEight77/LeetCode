# 13. Roman to Integer
## 解题思路
### map + 遍历
时间复杂度分析 O(n)，先用 map 记录每个字符对应的值，然后遍历，
我这里是判断 i 和 i - 1 的关系进行减法运算，每次减去前一个字符 * 2 的值，举例子
IV sum = 1 + 5 实际为 4 = 6 - 1 * 2
```
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
```
### 优化
由于给出的数据都是合法的罗马数字，即符合前一字符的值小于等于当前值，
所以单次遍历时候遇到 IV 时，判断当前下标对应值是否小于 i + 1 对应值，是的话，减去即可。 
