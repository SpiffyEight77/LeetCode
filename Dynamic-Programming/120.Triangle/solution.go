package main

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	var (
		l = len(triangle)
		f = make([][]int, l)
	)

	for i := 0; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			}
			f[i][j] = triangle[i][j]
		}
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			f[i][j] = min(f[i+1][j], f[i+1][j+1]) + triangle[i][j]
		}
	}
	return f[0][0]
}

func minimumTotalV2(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	var (
		l = len(triangle)
		f = make([][]int, l)
	)

	for i := 0; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			}
			f[i][j] = triangle[i][j]
		}
	}

	for i := 1; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				f[i][j] += f[i-1][j]
			} else if j == len(f[i-1]) {
				f[i][j] += f[i-1][j-1]
			} else {
				f[i][j] += min(f[i-1][j], f[i-1][j-1])
			}
		}
	}

	res := f[l-1][0]
	for i := 1; i < len(f[l-1]); i++ {
		res = min(res, f[l-1][i])
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
