package kata

type dir int

const (
	down dir = iota
	left
	up
	right
)

type location struct {
	x     int
	y     int
	dir   dir
	index int
}

var clueLoc = [16]location{
	{0, 0, down, 0},
	{0, 1, down, 1},
	{0, 2, down, 2},
	{0, 3, down, 3},
	{0, 3, left, 4},
	{1, 3, left, 5},
	{2, 3, left, 6},
	{3, 3, left, 7},
	{3, 3, up, 8},
	{3, 2, up, 9},
	{3, 1, up, 10},
	{3, 0, up, 11},
	{3, 0, right, 12},
	{2, 0, right, 13},
	{1, 0, right, 14},
	{0, 0, right, 15},
}

func findClueLoc(row, col int, lastIndex int, lastDir dir) location {
	for _, c := range clueLoc {
		if c.x == row && c.y == col && c.index > lastIndex && c.dir != lastDir {
			return c
		}
	}

	panic("invalid-row-col")
}

// clues helper
func getClues(row, col int) [4]location {
	w := findClueLoc(0, col, -1, right)
	x := findClueLoc(row, 3, w.index, w.dir)
	y := findClueLoc(3, col, x.index, x.dir)
	z := findClueLoc(row, 0, y.index, y.dir)
	locs := [4]location{
		w,
		x,
		y,
		z,
	}

	return locs
}

func scenes(row int, col int, d dir, board [4][4]int) int {
	sceene := 0
	maxScene := 0
	switch d {
	case down:
		for i := 0; i <= 3; i++ {
			if board[i][col] > maxScene {
				maxScene = board[i][col]
				sceene++
			}
		}
		break
	case left:
		for i := 3; i >= 0; i-- {
			if board[row][i] > maxScene {
				maxScene = board[row][i]
				sceene++
			}
		}
		break
	case up:
		for i := 3; i >= 0; i-- {
			if board[i][col] > maxScene {
				maxScene = board[i][col]
				sceene++
			}
		}
		break
	case right:
		for i := 0; i <= 3; i++ {
			if board[row][i] > maxScene {
				maxScene = board[row][i]
				sceene++
			}
		}
		break
	}

	return sceene
}

func violateSudokuRule(board [4][4]int, row, col int, num int) bool {
	for i := 0; i < 4; i++ {
		if board[row][i] == num {
			return true
		}
	}

	for j := 0; j < 4; j++ {
		if board[j][col] == num {
			return true
		}
	}

	return false
}

func violateSkyscraperRule(board [4][4]int, clues []int, row, col int) bool {
	clueLocations := getClues(row, col)
	for _, c := range clueLocations {
		view := scenes(c.x, c.y, c.dir, board)
		if clues[c.index] != 0 && view != clues[c.index] {
			return true
		}
	}
	return false
}

func violate(board [4][4]int, clues []int, row, col, num int) bool {
	if violateSudokuRule(board, row, col, num) {
		return true
	}

	if violateSkyscraperRule(board, clues, row, col) {
		return true
	}

	return false
}

func isCompleted(board [4][4]int) (is bool, row, col int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if board[i][j] == 0 {
				return false, i, j
			}
		}
	}

	return true, -1, -1
}

func solveSudoku(board [4][4]int) (bool, [4][4]int) {
	cmp, row, col := isCompleted(board)
	if cmp {
		return true, board
	}

	for i := 4; i > 0; i-- {
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

func solveSkyscraperSudoku(clues []int, board [4][4]int) (bool, [4][4]int) {
	cmp, row, col := isCompleted(board)
	if cmp {
		return true, board
	}

	for i := 4; i > 0; i-- {
		board[row][col] = i
		if !violate(board, clues, row, col, i) {
			ok, x := solveSkyscraperSudoku(clues, board)
			if ok {
				return ok, x
			}
		} else {
			board[row][col] = 0
		}
	}

	return false, board
}

func SolvePuzzle(clues []int) [][]int {
	// Your code

	board := [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	ok, b := solveSkyscraperSudoku(clues, board)
	if ok {
		x := make([][]int, 0)
		for i := 0; i < 4; i++ {
			x[i] = make([]int, 0)
		}

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				x[i] = append(x[i], b[i][j])
			}
		}

		return x
	}

	return nil
}
