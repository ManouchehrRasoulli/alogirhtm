package day7

import (
	"strings"
)

const (
	Start string = "S"
	Empty string = "."
	Split string = "^"
	Steam string = "|"
)

func Manifold(input string) int {
	var (
		field      [][]string
		updateFunc = func(field [][]string, x, y int, value string) {
			if x < 0 || x >= len(field) ||
				y < 0 || y >= len(field[x]) {
				return
			}

			field[x][y] = value
		}
		valueFunc = func(field [][]string, x, y int) string {
			if x < 0 || x >= len(field) ||
				y < 0 || y >= len(field[x]) {
				return Empty
			}
			return field[x][y]
		}
		lines      = strings.Split(input, "\n")
		splitCount = 0
	)

	for _, line := range lines {
		field = append(field, strings.Split(line, ""))
	}

	for i := 1; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if valueFunc(field, i-1, j) == Start {
				updateFunc(field, i, j, Steam)
				continue
			}
			if valueFunc(field, i, j) == Split {
				if valueFunc(field, i-1, j) == Steam {
					updateFunc(field, i, j-1, Steam)
					updateFunc(field, i, j+1, Steam)
				}
				continue
			}
			if valueFunc(field, i-1, j) == Steam && valueFunc(field, i, j) == Empty {
				updateFunc(field, i, j, Steam)
				continue
			}
		}
	}

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if valueFunc(field, i, j-1) == Steam &&
				valueFunc(field, i, j) == Split &&
				valueFunc(field, i, j+1) == Steam &&
				valueFunc(field, i-1, j) == Steam {
				splitCount++
			}
		}
	}

	return splitCount
}
