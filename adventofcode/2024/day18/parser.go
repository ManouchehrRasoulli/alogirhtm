package day18

import (
	"strconv"
	"strings"
)

func ReadPuzzle(data string, count int, row, col int) *Puzzle {
	lines := strings.Split(data, "\n")

	p := Puzzle{
		grid: make([][]rune, 0),
		row:  row,
		col:  col,
	}

	for i := 0; i <= row; i++ {
		line := make([]rune, 0)
		for j := 0; j <= col; j++ {
			line = append(line, freeSpace)
		}
		p.grid = append(p.grid, line)
	}

	for i, line := range lines {
		if i >= count {
			break
		}

		lineData := strings.Split(line, ",")
		x, _ := strconv.Atoi(lineData[0])
		y, _ := strconv.Atoi(lineData[1])

		p.grid[y][x] = usedSpace
	}

	return &p
}
