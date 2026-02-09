package day12

import (
	"strconv"
	"strings"
)

func rotate90(shp [3][3]rune) [3][3]rune {
	return [3][3]rune{
		{shp[2][0], shp[1][0], shp[0][0]},
		{shp[2][1], shp[1][1], shp[0][1]},
		{shp[2][2], shp[1][2], shp[0][2]},
	}
}

func rotate180(shp [3][3]rune) [3][3]rune {
	return rotate90(rotate90(shp))
}

func rotate270(shp [3][3]rune) [3][3]rune {
	return rotate180(rotate90(shp))
}

func rotate360(shp [3][3]rune) [3][3]rune {
	return rotate270(rotate90(shp))
}

func toGrid(in string) [][]rune {
	// format 10x6
	product := strings.Split(in, "x")
	x, y := toInt(product[0]), toInt(product[1])
	var grid [][]rune
	for i := 0; i < y; i++ {
		line := make([]rune, 0)
		for j := 0; j < x; j++ {
			line = append(line, empty)
		}
		grid = append(grid, line)
	}
	return grid
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func toIntArray(s string) []int {
	items := strings.Split(s, " ")
	result := make([]int, 0)
	for _, item := range items {
		item = strings.TrimSpace(item)
		if len(item) == 0 {
			continue
		}
		result = append(result, toInt(item))
	}
	return result
}

func shapeFitInGridAt(x, y int, grid [][]rune, shape [3][3]rune) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if y+j >= len(grid) {
				return false
			}
			if x+i >= len(grid[y+j]) {
				return false
			}
			if shape[i][j] != empty && grid[y+j][x+i] != empty {
				return false
			}
		}
	}
	return true
}

func drawShape(x, y int, grid [][]rune, shape [3][3]rune) bool {
	if !shapeFitInGridAt(x, y, grid, shape) {
		return false
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if shape[i][j] != empty {
				grid[y+j][x+i] = filled
			}
		}
	}

	return true
}

func clearShape(x, y int, grid [][]rune, shape [3][3]rune) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if shape[i][j] != empty {
				grid[y+j][x+i] = empty
			}
		}
	}
}

func gridKey(grid [][]rune) string {
	var b strings.Builder
	for _, row := range grid {
		b.WriteString(string(row))
	}
	return b.String()
}

func shapeWeight(s shape) int {
	count := 0
	for _, row := range s.indices[0] {
		for _, c := range row {
			if c == filled {
				count++
			}
		}
	}
	return count
}
