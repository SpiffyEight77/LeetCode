package leetcode

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	f := make([][]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			if i == 0 || j == 0 {
				f[i][j] = 1
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}
