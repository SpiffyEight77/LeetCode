package leetcode

func Merge(nums1 []int, m int, nums2 []int, n int)  {
  last, nums1Pointer, nums2Pointer := m + n - 1, m - 1, n - 1
  for nums1Pointer >= 0 && nums2Pointer >= 0 {
    if nums1[nums1Pointer] <= nums2[nums2Pointer] {
      nums1[last] = nums2[nums2Pointer]
      nums2Pointer--
    } else {
      nums1[last] = nums1[nums1Pointer]
      nums1Pointer--
    }
    last--
  }
  for nums2Pointer >= 0 {
    nums1[last] = nums2[nums2Pointer]
    last--
    nums2Pointer--
  }
}
