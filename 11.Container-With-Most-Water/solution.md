# 11.Container With Most Water
## 解题思路
### 嵌套循环
容易超时
### 双指针
时间复杂度分析 O(n)，左指针指向数组最左端，右指针指向数组最右端，
结束条件为左指针大于等于右指针。每次使用当前面积与最大面积比较，
左右指针分别向中心靠拢，最后返回最大面积
 ```
func MaxArea(height []int) int {
  left, right, maxArea := 0, len(height) - 1, 0
  for left < right {
    currentArea := min(height[left], height[right]) * (right - left)
    if height[left] > height[right] {
      right--
    } else {
      left++
    }
    maxArea = max(maxArea, currentArea)
  }
  return maxArea
}

func min(x, y int) int {
  if x > y {
    return y
  }
  return x
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}
 ```
