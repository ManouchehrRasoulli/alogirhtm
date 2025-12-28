package day9

import (
	"fmt"
	"strings"
)

func MaxRectangleInPolygon(input string) int {
	var (
		lines          = strings.Split(strings.TrimSpace(input), "\n")
		points         [][2]int
		maxRow, maxCol int
		minRow, minCol int = int(^uint(0) >> 1), int(^uint(0) >> 1)
		grid           [][]rune
		redTiles       [][2]int
	)

	for _, line := range lines {
		if line == "" {
			continue
		}
		var col, row int
		_, err := fmt.Sscanf(line, "%d,%d", &col, &row)
		if err != nil {
			panic(fmt.Sprintf("failed to parse line %q: %v", line, err))
		}

		maxRow = max(maxRow, row)
		minRow = min(minRow, row)
		maxCol = max(maxCol, col)
		minCol = min(minCol, col)

		points = append(points, [2]int{row, col})
	}

	height := maxRow - minRow + 1
	width := maxCol - minCol + 1

	grid = make([][]rune, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			grid[i][j] = '.'
		}
	}

	for _, point := range points {
		r := point[0] - minRow
		c := point[1] - minCol
		grid[r][c] = '#'
		redTiles = append(redTiles, [2]int{r, c})
	}

	drawLine := func(from, to [2]int) {
		r1, c1 := from[0]-minRow, from[1]-minCol
		r2, c2 := to[0]-minRow, to[1]-minCol

		if r1 == r2 {
			if c1 > c2 {
				c1, c2 = c2, c1
			}
			for c := c1; c <= c2; c++ {
				if grid[r1][c] == '.' {
					grid[r1][c] = 'X'
				}
			}
		} else if c1 == c2 {
			if r1 > r2 {
				r1, r2 = r2, r1
			}
			for r := r1; r <= r2; r++ {
				if grid[r][c1] == '.' {
					grid[r][c1] = 'X'
				}
			}
		}
	}

	for i := 1; i < len(points); i++ {
		drawLine(points[i-1], points[i])
	}
	drawLine(points[len(points)-1], points[0])

	// Scanline fill
	for r := 0; r < height; r++ {
		inside := false
		for c := 0; c < width; c++ {
			if grid[r][c] == '#' || grid[r][c] == 'X' {
				// Check if next cell is empty (boundary edge)
				if c+1 < width && grid[r][c+1] == '.' {
					inside = !inside
				} else if c+1 >= width {
					inside = !inside
				}
			} else if inside {
				grid[r][c] = 'X'
			}
		}
	}

	isValidRectangle := func(r1, c1, r2, c2 int) bool {
		minR, maxR := min(r1, r2), max(r1, r2)
		minC, maxC := min(c1, c2), max(c1, c2)

		for r := minR; r <= maxR; r++ {
			for c := minC; c <= maxC; c++ {
				if grid[r][c] != '#' && grid[r][c] != 'X' {
					return false
				}
			}
		}
		return true
	}

	maxArea := 0
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			r1, c1 := redTiles[i][0], redTiles[i][1]
			r2, c2 := redTiles[j][0], redTiles[j][1]

			if r1 != r2 && c1 != c2 {
				if isValidRectangle(r1, c1, r2, c2) {
					height := abs(r2-r1) + 1
					width := abs(c2-c1) + 1
					area := height * width

					if area > maxArea {
						maxArea = area
					}
				}
			}
		}
	}

	return maxArea
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
