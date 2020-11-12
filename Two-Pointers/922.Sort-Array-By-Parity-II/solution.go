package leetcode

func sortArrayByParityII(A []int) []int {
	if len(A) < 2 {
		return nil
	}
	odd := make([]int, 0)
	even := make([]int, 0)
	res := make([]int, 0)

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			even = append(even, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}

	for i := 0; i < len(A)/2; i++ {
		res = append(res, even[i], odd[i])
	}

	return res
}
