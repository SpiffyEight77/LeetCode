# 1. Two Sum
## 解题思路
### 嵌套循环
时间复杂度分析 O(n^2) n * (n + 1)，直接两层循环得出结果
```
```

### hash map
时间复杂度分析 O(n)，map[value] = requireIndex，记录当前值对应的下标
```
func twoSum(nums []int, target int) []int {
  requireIndexMap := make(map[int]int)
  for currentIndex, value := range nums {
    if requireIndex, ok := requireIndexMap[target - value]; ok {
      return []int{requireIndex, currentIndex}
    } else {
      requireIndexMap[value] = currentIndex
    }
  }
  return nil
}
```

### 双指针
时间复杂度分析 O(n^2) n * (n + 1)，思路跟嵌套循环基本一致，左指针先不动，
右指针遍历，当右指针遍历到数组末端时，左指针+1，右指针等于左指针+1，然后遍历
```

```

