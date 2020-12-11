package leetcode

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}

	index := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[index] {
				temp := board[i][j]
				board[i][j] = ' '
				res := dfs(i, j, &board, &word, index+1)
				if res == true {
					return true
				}
				board[i][j] = temp
			}
		}
	}
	return false
}

func dfs(i, j int, board *[][]byte, word *string, index int) bool {
	if len(*word) == index {
		return true
	}

	direction := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for k := 0; k < 4; k++ {
		x := i + direction[k][0]
		y := j + direction[k][1]

		if x < len(*board) && x >= 0 && y < len((*board)[0]) && y >= 0 && (*board)[x][y] == (*word)[index] {
			temp := (*board)[x][y]
			(*board)[x][y] = ' '
			res := dfs(x, y, board, word, index+1)
			if res == true {
				return true
			}
			(*board)[x][y] = temp
		}
	}
	return false
}
