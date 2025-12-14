package sudoko

func violateSudokuRule(board *[9][9]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return true
		}
	}

	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return true
		}
	}

	gridRow := (row / 3) * 3
	gridCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[gridRow+i][gridCol+j] == num {
				return true
			}
		}
	}

	return false
}

func nextRowCol(row, col int) (int, int) {
	col++
	if col >= 9 {
		row++
		col = 0
	}
	if row >= 9 {
		return -1, -1
	}
	return row, col
}

func Solve(board *[9][9]int, row, col int) bool {
	if row == -1 && col == -1 {
		return true
	}

	if (*board)[row][col] != 0 {
		nextRow, nextCol := nextRowCol(row, col)
		return Solve(board, nextRow, nextCol)
	}

	for i := 1; i <= 9; i++ {
		if violateSudokuRule(board, row, col, i) {
			continue
		}

		(*board)[row][col] = i
		nextRow, nextCol := nextRowCol(row, col)
		if Solve(board, nextRow, nextCol) {
			return true
		}
		(*board)[row][col] = 0
	}

	return false
}
