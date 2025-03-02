package _1_path_sum

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

func MinPathSum(input [][]int) int {
	minMatrix := make([][]int, 0)
	for i := 0; i < len(input); i++ {
		minMatrix = append(minMatrix, make([]int, 0))
		for j := 0; j < len(input[i]); j++ {
			minMatrix[i] = append(minMatrix[i], math.MaxInt64)
		}
	}

	minPathSum(input, minMatrix, len(input)-1, len(input)-1)

	return minMatrix[0][0]
}

func minPathSum(input [][]int, minPaths [][]int, row, col int) {
	if row == -1 || col == -1 {
		return
	}

	if row == len(input)-1 && col == len(input)-1 {
		// -- first time only --
		minPaths[row][col] = input[row][col]

		minPathSum(input, minPaths, row, col-1)
		minPathSum(input, minPaths, row-1, col)

		return
	}

	if row == len(input)-1 {
		rightMin := minPaths[row][col+1]
		if rightMin == math.MaxInt64 {
			return // this could not happen !!
		}

		minPaths[row][col] = input[row][col] + rightMin

		minPathSum(input, minPaths, row, col-1)
		minPathSum(input, minPaths, row-1, col)

		return
	}

	if col == len(input)-1 {
		downMin := minPaths[row+1][col]
		if downMin == math.MaxInt64 {
			return // this could not happen !!
		}

		minPaths[row][col] = input[row][col] + downMin

		minPathSum(input, minPaths, row, col-1)
		minPathSum(input, minPaths, row-1, col)

		return
	}

	rightMin := minPaths[row][col+1]
	downMin := minPaths[row+1][col]

	if rightMin == math.MaxInt64 || downMin == math.MaxInt64 {
		return
	}

	minPaths[row][col] = input[row][col] + min(rightMin, downMin)

	minPathSum(input, minPaths, row, col-1)
	minPathSum(input, minPaths, row-1, col)

	return
}

var (
	//go:embed input.txt
	fileContent string

	matrix [][]int
)

func init() {
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		items := strings.Split(line, ",")
		row := make([]int, len(items))
		for i, item := range items {
			v, _ := strconv.ParseInt(item, 10, 64)
			row[i] = int(v)
		}
		matrix = append(matrix, row)
	}
}
