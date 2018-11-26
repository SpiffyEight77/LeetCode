package main

//v1 O(n) O(n)
func sortArrayByParity(A []int) []int {
	even := make([]int, 0, len(A))
	odd := make([]int, 0, len(A))

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			even = append(even, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}
	even = append(even, odd...)
	return even
}

//v2 O(n) O(n)
func sortArrayByParity(A []int) []int {
	tmp := 0
	for i, j := 0, 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			tmp = A[j]
			A[j] = A[i]
			A[i] = tmp
			j++
		}
	}
	return A
}
