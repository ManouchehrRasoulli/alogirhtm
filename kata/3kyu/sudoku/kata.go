package kata

func violateSudokuRule(board [9][9]int, row, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return true
		}
	}

	for j := 0; j < 9; j++ {
		if board[j][col] == num {
			return true
		}
	}

	return false
}

func isCompleted(board [9][9]int) (is bool, row, col int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false, i, j
			}
		}
	}

	return true, -1, -1
}

func solveSudoku(board [9][9]int) (bool, [9][9]int) {
	cmp, row, col := isCompleted(board)
	if cmp {
		return true, board
	}

	for i := 1; i <= 9; i++ {
		if !violateSudokuRule(board, row, col, i) {
			board[row][col] = i
			ok, x := solveSudoku(board)
			if ok {
				return ok, x
			}
		} else {
			board[row][col] = 0
		}
	}

	return false, board
}
