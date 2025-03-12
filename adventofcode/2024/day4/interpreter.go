package day4

import (
	"log"
	"strings"
)

type Direction struct {
	x int
	y int
}

var (
	Up = Direction{
		x: -1,
		y: 0,
	}
	Down = Direction{
		x: 1,
		y: 0,
	}
	Left = Direction{
		x: 0,
		y: -1,
	}
	Right = Direction{
		x: 0,
		y: 1,
	}
	TopLeft = Direction{
		x: -1,
		y: -1,
	}
	TopRight = Direction{
		x: -1,
		y: 1,
	}
	BottomLeft = Direction{
		x: 1,
		y: -1,
	}
	BottomRight = Direction{
		x: 1,
		y: 1,
	}

	AlDirections = []Direction{Up, Down, Left, Right, TopLeft, TopRight, BottomLeft, BottomRight}
)

func ReadUpToXInDirection(x, y int, direction Direction, n int, matrix [][]string) []string {
	if n == 0 {
		return []string{}
	}

	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[x]) {
		return []string{}
	}

	return append([]string{matrix[x][y]},
		ReadUpToXInDirection(x+direction.x, y+direction.y, direction, n-1, matrix)...)
}

const (
	LookingFor       = "XMAS"
	RevereLookingFor = "SAMX"
)

func Neighbors(matrix [][]string, x, y int, n int) [][]string {
	ngs := make([][]string, 0)

	for _, direction := range AlDirections {
		ngData := ReadUpToXInDirection(x, y, direction, n, matrix)
		if len(ngData) < n {
			continue
		}

		ngs = append(ngs, ngData)
	}

	return ngs
}

func FindPattern(matrix [][]string) int {
	count := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "S" {
				ngs := Neighbors(matrix, i, j, 4)
				for _, ng := range ngs {
					ngString := strings.Join(ng, "")
					if ngString == RevereLookingFor {
						log.Println(i, j, ngString)
						count++
					}
				}
			}
		}
	}

	return count
}
