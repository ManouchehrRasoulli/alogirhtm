package kata

import "testing"

func TestViolateSudokuRule(t *testing.T) {
	var boar = [4][4]int{
		{2, 1, 4, 3},
		{3, 0, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4},
	}

	t.Log(violateSudokuRule(boar, 1, 1, 3))
	t.Log(violateSudokuRule(boar, 1, 1, 4))
}

func TestViolateSkyscraperRule(t *testing.T) {
	var clues = []int{
		0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 1, 0, 2}

	var boar = [4][4]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4},
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if violateSkyscraperRule(boar, clues, i, j) {
				t.Error(i, j)
			}
			x := boar[i][j]
			boar[i][j] = 0
			if violateSudokuRule(boar, i, j, x) {
				t.Error(i, j, "sudoko !")
			}
			boar[i][j] = x
		}
	}
}

func TestSolveSudoku(t *testing.T) {
	board := [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	ok, b := solveSudoku(board)
	t.Log(ok)

	for i := 0; i < 4; i++ {
		t.Log(b[i])
	}
}

func TestSceene(t *testing.T) {
	var boar = [4][4]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4},
	}

	t.Log(scenes(0, 0, right, boar))
	t.Log(scenes(0, 3, left, boar))
	t.Log(scenes(0, 3, down, boar))
	t.Log(scenes(3, 3, up, boar))
	t.Log(scenes(3, 2, up, boar))
	t.Log(scenes(2, 0, right, boar))
}

func TestFindClueLocation(t *testing.T) {
	var clues = []int{
		0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 1, 0, 2}

	t.Log(SolvePuzzle(clues))
}
