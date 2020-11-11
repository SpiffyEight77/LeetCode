package leetcode

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		m = len(grid)
		n = len(grid[0])
	)

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
