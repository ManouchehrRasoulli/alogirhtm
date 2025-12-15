package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func ManifoldDimensions(input string) int {
	var (
		field    [][]string
		isNumber = func(s string) bool {
			_, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return false
			}
			return true
		}
		addFunc = func(field [][]string, x, y int, value string) {
			if x < 0 || x >= len(field) ||
				y < 0 || y >= len(field[x]) {
				return
			}

			currentValue := field[x][y]
			currentDigit, _ := strconv.ParseInt(currentValue, 10, 64)
			addDigit, _ := strconv.ParseInt(value, 10, 64)

			field[x][y] = fmt.Sprintf("%d", currentDigit+addDigit)
		}
		valueFunc = func(field [][]string, x, y int) string {
			if x < 0 || x >= len(field) ||
				y < 0 || y >= len(field[x]) {
				return Empty
			}
			return field[x][y]
		}
		lines          = strings.Split(input, "\n")
		dimensionCount = 0
	)

	for _, line := range lines {
		field = append(field, strings.Split(line, ""))
	}

	for i := 1; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if valueFunc(field, i-1, j) == Start {
				addFunc(field, i, j, "1")
			}
			if valueFunc(field, i, j) == Split &&
				isNumber(valueFunc(field, i-1, j)) {
				addFunc(field, i, j-1, valueFunc(field, i-1, j))
				addFunc(field, i, j+1, valueFunc(field, i-1, j))
			}
			if isNumber(valueFunc(field, i-1, j)) &&
				valueFunc(field, i, j) != Split {
				addFunc(field, i, j, valueFunc(field, i-1, j))
			}
		}
	}

	for _, row := range field[len(field)-1] {
		num, _ := strconv.Atoi(row)
		dimensionCount += num
	}

	return dimensionCount
}
