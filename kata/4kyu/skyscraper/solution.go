package kata

import "fmt"

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

func findClueLoc(row, col int, lastIndex int) location {
	for _, c := range clueLoc {
		if c.x == row && c.y == col && c.index > lastIndex {
			return c
		}
	}

	panic("invalid-row-col")
}

// clues helper
func getClues(row, col int) [4]location {
	w := findClueLoc(0, col, -1)
	x := findClueLoc(row, 3, w.index)
	y := findClueLoc(3, col, x.index)
	z := findClueLoc(row, 0, y.index)
	locs := [4]location{
		w,
		x,
		y,
		z,
	}

	return locs
}

func scenes(row, col, d dir, board [4][4]int) int {
	sceene := 0
	maxScene := 0
	switch d {
	case down:
		for i := 0; i < 3; i++ {
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
		for i := 0; i < 3; i++ {
			if board[row][i] > maxScene {
				maxScene = board[row][i]
				sceene++
			}
		}
		break
	}

	return sceene
}

func matchClues(board [4][4]int, clues []int, row, col int) bool {
	clueLocations := getClues(row, col)
	for _, c := range clueLocations {
		fmt.Println(c, clues[c.index])
	}
	panic("not implemented")
}

func SolvePuzzle(clues []int) [][]int {
	// Your code

	board := [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	fmt.Println(board[1][2])

	panic("not implemented yet")
}
