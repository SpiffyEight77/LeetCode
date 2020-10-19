package leetcode

type Point struct {
	X, Y int
}

func maxDistance(grid [][]int) int {
	queue := make([]*Point, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, &Point{i, j})
			}
		}
	}
	if len(queue) == 0 || len(queue) == len(grid)*len(grid) {
		return -1
	}

	d := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var p *Point
	for len(queue) > 0 {
		p = queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := p.X + d[i][0]
			y := p.Y + d[i][1]
			if x >= len(grid) || x < 0 || y >= len(grid) || y < 0 || grid[x][y] != 0 {
				continue
			}
			grid[x][y] = grid[p.X][p.Y] + 1
			queue = append(queue, &Point{x, y})
		}
	}
	return grid[p.X][p.Y] - 1
}
