package day9

import (
	"fmt"
	"strings"
)

func MaxRectangleInPolygon(input string) int {
	var (
		lines              = strings.Split(strings.TrimSpace(input), "\n")
		points             [][2]int
		polygon            [][2][2]int
		pointOnSegmentFunc = func(p, a, b [2]int) bool {
			px, py := p[0], p[1]
			ax, ay := a[0], a[1]
			bx, by := b[0], b[1]
			cross := (px-ax)*(by-ay) - (py-ay)*(bx-ax)
			if cross != 0 {
				return false
			}
			minX, maxX := ax, bx
			minY, maxY := ay, by
			if bx < ax {
				minX, maxX = bx, ax
			}
			if by < ay {
				minY, maxY = by, ay
			}
			return px >= minX && px <= maxX && py >= minY && py <= maxY
		}
		rayCrossSegmentFunc = func(p, a, b [2]int) bool {
			px, py := float64(p[0]), float64(p[1])
			ax, ay := float64(a[0]), float64(a[1])
			bx, by := float64(b[0]), float64(b[1])
			if ay > by {
				ax, bx = bx, ax
				ay, by = by, ay
			}
			if ay == by {
				return false
			}
			if py <= ay || py > by {
				return false
			}
			intersectX := ax + (bx-ax)*(py-ay)/(by-ay)
			return px < intersectX
		}
		pointIsInPolygonFunc = func(p [2]int) bool {
			for _, edge := range polygon {
				if pointOnSegmentFunc(p, edge[0], edge[1]) {
					return true
				}
			}
			count := 0
			for _, edge := range polygon {
				if rayCrossSegmentFunc(p, edge[0], edge[1]) {
					count++
				}
			}
			return count%2 == 1
		}
		pointsAreInPolygonFunc = func(points [][2]int) bool {
			for _, point := range points {
				if !pointIsInPolygonFunc(point) {
					return false
				}
			}
			return true
		}
		rectangleFromDiagonalFunc = func(p1, p2 [2]int) [][2]int {
			x1, y1 := p1[0], p1[1]
			x2, y2 := p2[0], p2[1]
			return [][2]int{{x1, y1}, {x1, y2}, {x2, y2}, {x2, y1}}
		}
		abs = func(a int) int {
			if a < 0 {
				return -a
			}
			return a
		}
		linePointsFunc = func(p1, p2 [2]int) [][2]int {
			x1, y1 := p1[0], p1[1]
			x2, y2 := p2[0], p2[1]
			linePoints := make([][2]int, 0)
			dx := abs(x2 - x1)
			dy := abs(y2 - y1)
			sx := 1
			if x1 > x2 {
				sx = -1
			}
			sy := 1
			if y1 > y2 {
				sy = -1
			}
			err := dx - dy
			for {
				linePoints = append(linePoints, [2]int{x1, y1})
				if x1 == x2 && y1 == y2 {
					break
				}
				e2 := 2 * err
				if e2 > -dy {
					err -= dy
					x1 += sx
				}
				if e2 < dx {
					err += dx
					y1 += sy
				}
			}
			return linePoints
		}
		rectangleAreaFunc = func(p1, p2 [2]int) int {
			width := abs(p2[0]-p1[0]) + 1
			height := abs(p2[1]-p1[1]) + 1
			return width * height
		}
		maxSizeInsidePolygon = 0
	)

	for _, line := range lines {
		x, y := 0, 0
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		points = append(points, [2]int{x, y})
	}

	for i := 1; i < len(points); i++ {
		polygon = append(polygon, [2][2]int{points[i-1], points[i]})
	}
	polygon = append(polygon, [2][2]int{points[len(points)-1], points[0]})

	for x := 0; x < len(points); x++ {
	bLoop:
		for y := x + 1; y < len(points); y++ {
			a, b := points[x], points[y]
			rectanglePoints := rectangleFromDiagonalFunc(a, b)
			for i := 1; i < len(rectanglePoints); i++ {
				linePoints := linePointsFunc(rectanglePoints[i], rectanglePoints[i-1])
				if !pointsAreInPolygonFunc(linePoints) {
					continue bLoop
				}
			}
			linePoints := linePointsFunc(rectanglePoints[len(rectanglePoints)-1], rectanglePoints[0])
			if !pointsAreInPolygonFunc(linePoints) {
				continue bLoop
			}
			maxSizeInsidePolygon = max(maxSizeInsidePolygon, rectangleAreaFunc(a, b))
		}
	}

	return maxSizeInsidePolygon
}
