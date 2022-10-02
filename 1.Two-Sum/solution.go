package leetcode

func TwoSum(nums []int, target int) []int {
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


func twoSumForLoop(nums []int, target int) []int {
  for left:= 0; left < len(nums); left++ {
    for right:= left + 1; right < len(nums); right++ {
      if (nums[left] + nums[right] == target) {
        return []int{left, right}
      }
    }
  }
  return nil
}


func twoSumTwoPointers(nums []int, target int) []int {
  left,right := 0, 1
  for {
    if (nums[left] + nums[right] == target) {
      return []int{left, right}
    }
    right++
    if (right == len(nums)) {
      left++
      right = left + 1 
    }
  }
}
