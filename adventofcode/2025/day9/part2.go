package day9

import (
	"fmt"
	"strings"
)

func MaxRectangleInPolygon(input string) int {
	var (
		err            error
		grid           [][]rune
		maxRow, maxCol int
		minRow, minCol int
		lines          = strings.Split(input, "\n")
		points         [][2]int
		printGridFunc  = func() {
			for _, row := range grid {
				fmt.Println(string(row))
			}
		}
		drawLineFunc = func(from, to [2]int) {
			var (
				horizontal bool
				vertical   bool
			)

			if from[0] == to[0] {
				horizontal = true
			}
			if from[1] == to[1] {
				vertical = true
			}

			if !vertical && !horizontal {
				fmt.Printf("could not draw line from %+v to %+v", from, to)
				return
			}

			if horizontal {
				fromY, toY := from[1]-minCol, to[1]-minCol
				if fromY > toY {
					fromY, toY = toY, fromY
				}
				for i := fromY; i <= toY; i++ {
					grid[from[0]][i] = 'X'
				}
			}

			if vertical {
				fromX, toX := from[0]-minRow, to[0]-minRow
				if fromX > toX {
					fromX, toX = toX, fromX
				}
				for i := fromX; i <= toX; i++ {
					grid[i][from[1]] = 'X'
				}
			}
		}
	)

	for _, line := range lines {
		col, row := 0, 0
		_, err = fmt.Sscanf(line, "%d,%d", &col, &row)
		if err != nil {
			panic(err)
		}
		maxCol, maxRow = max(maxCol, col), max(maxRow, row)
		minCol, minRow = min(minCol, col), min(minRow, row)
		points = append(points, [2]int{row, col})
	}

	maxRow = maxRow - minRow + 1
	maxCol = maxCol - minCol + 1

	fmt.Println(maxRow * maxCol)

	grid = make([][]rune, maxRow)
	for i := 0; i < maxRow; i++ {
		grid[i] = make([]rune, maxCol)
		for j := 0; j < maxCol; j++ {
			grid[i][j] = '.'
		}
	}

	for _, point := range points {
		grid[point[0]-minRow][point[1]-minCol] = '#'
	}

	printGridFunc()
	fmt.Println("************************")
	for i := 1; i < len(points); i++ {
		drawLineFunc(points[i-1], points[i])
	}
	// close the circuit
	drawLineFunc(points[0], points[len(points)-1])

	printGridFunc()

	return 0
}
