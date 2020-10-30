package leetcode

type position struct {
	x, y int
}

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	var res int
	queue := make([]position, 0)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, position{i, j})
				res += 4
			}
		}
	}

	if len(queue) == 0 {
		return 0
	}

	direction := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := p.x + direction[i][0]
			y := p.y + direction[i][1]
			if x >= 0 && x < row && y >= 0 && y <= col && grid[x][y] == 1 {
				res -= 1
			}
		}
	}
	return res
}
