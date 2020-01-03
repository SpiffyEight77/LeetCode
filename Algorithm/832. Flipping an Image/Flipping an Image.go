package main

//v1 O(N*N)
func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A)/2; j++ {
			t := A[i][j]
			A[i][j] = A[i][len(A)-j-1]
			A[i][len(A)-j-1] = t
		}
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			A[i][j] = A[i][j] ^ 1
		}
	}
	return A
}
