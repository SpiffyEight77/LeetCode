package leetcode

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
