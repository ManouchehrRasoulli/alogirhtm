package day12

import (
	"strings"
)

func FitShapes(input string) int {
	var (
		lines   = strings.Split(input, "\n")
		shapes  []shape
		regions []region
	)

	for i := 0; i < len(lines); {
		// nothings
		if len(lines[i]) == 0 {
			i++
			continue
		}

		// read in shapes
		if len(shapes) < 6 {
			currentShape := [3][3]rune{}
			for j := 1; j <= 3; j++ {
				ln := []rune(lines[i+j])
				currentShape[j-1] = [3]rune{ln[0], ln[1], ln[2]}
			}
			shapes = append(shapes, newShape(currentShape))
			i += 5
			continue
		}

		// read in placements
		part := strings.Split(lines[i], ":")
		grid := toGrid(part[0])
		quantities := toIntArray(part[1])
		regions = append(regions, newRegion(grid, shapes, quantities))
		i++
	}

	return 0
}
