package leetcode

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 0 {
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return 0
	}

	if grid[i][j] == '1' {
		grid[i][j] = 0
		return dfs(grid, i+1, j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1) + 1
	}
	return 0
}
