package leetcode

func sortArrayByParityII(A []int) []int {
	if len(A) == 0 {
		return nil
	}

	for i, j := 0, 1; i < len(A); i += 2 {
		if A[i]%2 == 1 {
			for A[j]%2 == 1 && j < len(A) {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
