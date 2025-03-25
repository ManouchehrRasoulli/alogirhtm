package day20

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
)

func ReadPuzzle(data string) *Puzzle {
	lines := strings.Split(data, "\n")

	p := Puzzle{
		grid:  make([][]rune, 0),
		row:   0,
		col:   0,
		start: nil,
		end:   nil,
		path:  make([]*helper.Location, 0),
	}

	for i, line := range lines {
		lineData := make([]rune, 0)
		for j, char := range line {
			lineData = append(lineData, char)
			p.col = max(p.col, j)

			switch char {
			case startLocation:
				p.start = helper.NewLocation(i, j)
			case endLocation:
				p.end = helper.NewLocation(i, j)
			}
		}

		p.grid = append(p.grid, lineData)
		p.row = max(p.row, i)
	}

	return &p
}
