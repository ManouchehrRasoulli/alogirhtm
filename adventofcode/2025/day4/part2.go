package day3

import (
	"strings"
)

func markAccessAbleRoles(field [][]string) int {
	var (
		directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		fieldValue = func(x, y int) string {
			if x < 0 || x >= len(field) || y < 0 || y >= len(field[x]) {
				return empty
			}

			return field[x][y]
		}
		neighborRolesCount = func(x, y int) int {
			count := 0
			for _, dir := range directions {
				v := fieldValue(x+dir[0], y+dir[1])
				if v == role || v == accessAble {
					count++
				}
			}
			return count
		}
		accessCount = 0
	)

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if fieldValue(i, j) == role &&
				neighborRolesCount(i, j) < 4 {
				accessCount++
				field[i][j] = accessAble
			}
		}
	}

	return accessCount
}

func RemoveAccessAbleRoles(input string) int {
	var (
		field                = make([][]string, 0)
		lines                = strings.Split(input, "\n")
		clearAccessAbleRoles = func(field [][]string) {
			for i := 0; i < len(field); i++ {
				for j := 0; j < len(field[i]); j++ {
					if field[i][j] == accessAble {
						field[i][j] = empty
					}
				}
			}
		}
		accessCount = 0
	)
	for _, line := range lines {
		field = append(field, strings.Split(line, ""))
	}

	currentAccessCount := markAccessAbleRoles(field)
	accessCount = currentAccessCount
	for currentAccessCount > 0 {
		clearAccessAbleRoles(field)
		currentAccessCount = markAccessAbleRoles(field)
		accessCount += currentAccessCount
	}

	return accessCount
}
