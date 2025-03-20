package day16

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
)

func MazeFromData(data string) *Maze {
	var (
		maze = Maze{
			grid:  make([][]rune, 0),
			start: nil,
			end:   nil,
		}
		lines = strings.Split(data, "\n")
	)

	for i, row := range lines {
		lineData := make([]rune, 0)
		for j, char := range row {
			if char == start {
				maze.start = helper.NewLocation(i, j)
			}
			if char == end {
				maze.end = helper.NewLocation(i, j)
			}

			lineData = append(lineData, char)
		}

		maze.grid = append(maze.grid, lineData)
	}

	return &maze
}
