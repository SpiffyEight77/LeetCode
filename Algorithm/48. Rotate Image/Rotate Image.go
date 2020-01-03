package main

import "fmt"

//v1
func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		for j := i; j < l-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[l-j-1][i]
			matrix[l-j-1][i] = matrix[l-i-1][l-j-1]
			matrix[l-i-1][l-j-1] = matrix[j][l-i-1]
			matrix[j][l-i-1] = tmp
 		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}
