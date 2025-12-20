package day9

import (
	"fmt"
	"strings"
)

func MaxRectangleInPolygon(input string) int {
	var (
		lines       = strings.Split(input, "\n")
		maxValue    = 0
		polygon     [][2]int
		isInPolygon = func(point [2]int, polygon [][2]int) bool {
			panic("not implemented")
		}
		rectangle = func(a, b [2]int) [4][2]int {
			panic("not implemented")
		}
	)

	_ = isInPolygon
	_ = rectangle

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
			// todo
			//		1 - generate rectangle
			//		2 - check points are inside the polygon
		}
	}

	return maxValue
}
