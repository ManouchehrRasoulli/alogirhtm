package day9

import (
	"fmt"
	"strings"
)

func MaxRectangle(input string) int {
	var (
		lines         = strings.Split(input, "\n")
		maxValue      = 0
		polygon       [][2]int
		rectangleSize = func(a, b [2]int) int {
			ax := a[0] - b[0]
			bx := a[1] - b[1]
			if ax < 0 {
				ax = -ax
			}
			if bx < 0 {
				bx = -bx
			}
			return (ax + 1) * (bx + 1)
		}
	)

	for _, line := range lines {
		var a, b int
		_, err := fmt.Sscanf(line, "%d,%d", &a, &b)
		if err != nil {
			panic(err)
		}
		polygon = append(polygon, [2]int{a, b})
	}

	for i := 0; i < len(polygon); i++ {
		for j := i + 1; j < len(polygon); j++ {
			maxValue = max(maxValue, rectangleSize(polygon[i], polygon[j]))
		}
	}

	return maxValue
}
