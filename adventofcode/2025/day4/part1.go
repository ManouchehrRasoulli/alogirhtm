package day3

import (
	"strings"
)

const (
	empty      string = "."
	role       string = "@"
	accessAble string = "x"
)

func AccessAbleRoles(input string) int {
	var (
		field      = make([][]string, 0)
		lines      = strings.Split(input, "\n")
		directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		fieldValue = func(x, y int) string {
			if x < 0 || x >= len(field) {
				return empty
			}
			if y < 0 || y >= len(field[x]) {
				return empty
			}

			return field[x][y]
		}
		neighborRolesCount = func(x, y int) int {
			count := 0
			for _, dir := range directions {
				v := fieldValue(x+dir[0], y+dir[1])
				if v == role {
					count++
				}
			}
			return count
		}
		accessCount = 0
	)
	for _, line := range lines {
		field = append(field, strings.Split(line, ""))
	}

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if fieldValue(i, j) == role &&
				neighborRolesCount(i, j) < 4 {
				accessCount++
			}
		}
	}

	return accessCount
}
