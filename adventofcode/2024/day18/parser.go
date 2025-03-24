package day18

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strconv"
	"strings"
)

func ReadLocations(data string) []*helper.Location {
	var (
		location = make([]*helper.Location, 0)
	)

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		lineData := strings.Split(line, ",")
		y, _ := strconv.Atoi(lineData[0])
		x, _ := strconv.Atoi(lineData[1])

		location = append(location, helper.NewLocation(x, y))
	}

	return location
}

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
		y, _ := strconv.Atoi(lineData[0])
		x, _ := strconv.Atoi(lineData[1])

		p.grid[x][y] = usedSpace
	}

	return &p
}
