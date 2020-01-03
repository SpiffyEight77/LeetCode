package main

import "fmt"

//v1
//func nextGreaterElement(findNums []int, nums []int) []int {
//	res := make([]int, 0, len(findNums))
//	dict := make(map[int]int)
//	for k, v := range nums {
//		dict[v] = k
//	}
//
//	for i := 0; i < len(findNums); i++ {
//		j := dict[findNums[i]] + 1
//		for {
//			if j == len(nums) {
//				res = append(res, -1)
//				break
//			}
//			if nums[j] > findNums[i] {
//				res = append(res, nums[j])
//				break
//			}
//			j++
//		}
//	}
//	return res
//}

//v2
func nextGreaterElement(findNums []int, nums []int) []int {

}

func main() {
	findNums := []int{2, 4}
	nums := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(findNums, nums))
}
