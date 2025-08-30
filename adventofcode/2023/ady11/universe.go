package ady11

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"strings"
)

const (
	galaxy = '#'
	empty  = '.'
)

func ReadUniverse(input string) [][]rune {
	universe := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		universe = append(universe, []rune(line))
	}

	return universe
}

func ExpandUniverse(universe [][]rune) [][]rune {
	expandedUniverse := make([][]rune, 0)

	rowNeedToExpand := func(row int) bool {
		for j := 0; j < len(universe[row]); j++ {
			if universe[row][j] == galaxy {
				return false
			}
		}
		return true
	}

	colNeedToExpand := func(col int) bool {
		for i := 0; i < len(universe); i++ {
			if universe[i][col] == galaxy {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(universe); i++ {
		if rowNeedToExpand(i) { // append twice
			rowCopy := make([]rune, len(universe[i]))
			copy(rowCopy, universe[i])
			expandedUniverse = append(expandedUniverse, rowCopy)
		}
		rowCopy := make([]rune, len(universe[i]))
		copy(rowCopy, universe[i])
		expandedUniverse = append(expandedUniverse, rowCopy)
	}

	colExpandedCount := 0
	for j := 0; j < len(universe[0]); j++ {
		if colNeedToExpand(j) {
			colExpandedCount++
			// for each row insert an empty space in j field
			for i := 0; i < len(expandedUniverse); i++ {
				expandedUniverse[i] = helper.Insert(expandedUniverse[i], j+colExpandedCount, empty)
			}
		}
	}

	return expandedUniverse
}

func PrintUniverse(universe [][]rune) {
	for _, line := range universe {
		fmt.Println(string(line))
	}
}
