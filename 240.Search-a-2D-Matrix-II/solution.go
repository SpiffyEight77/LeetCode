package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
  row, col := len(matrix), len(matrix[0])
  x, y := 0, col - 1
  for x < row && y >= 0 {
    if matrix[x][y] == target {
      return true
    }
    if target > matrix[x][y] {
      x++ 
    } else {
      y--
    }
  }
  return false    
}

