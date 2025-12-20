package day9

import (
	"fmt"
	"strings"
)

func MaxRectangle(input string) int {
	var (
		lines    = strings.Split(input, "\n")
		maxValue = 0
		points   [][2]int
		distance = func(a, b [2]int) int {
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
		points = append(points, [2]int{a, b})
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			maxValue = max(maxValue, distance(points[i], points[j]))
		}
	}

	return maxValue
}
