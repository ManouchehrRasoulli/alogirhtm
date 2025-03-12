package day4

import (
	"log"
	"regexp"
	"strings"
)

type Direction struct {
	x int
	y int
}

var (
	Top = Direction{
		x: -1,
		y: 0,
	}
	Bottom = Direction{
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

	AlDirections = []Direction{TopLeft, Top, TopRight, Left, Right, BottomLeft, Bottom, BottomRight}
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

func FindPatternConstant(matrix [][]string) int {
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

func Matrix1Snippet(matrix [][]string, x, y int) []string {
	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[x]) {
		return []string{}
	}

	return []string{
		ReadUpToXInDirection(TopLeft.x+x, TopLeft.y+y, TopLeft, 1, matrix)[0],
		ReadUpToXInDirection(Top.x+x, Top.y+y, Top, 1, matrix)[0],
		ReadUpToXInDirection(TopRight.x+x, TopRight.y+y, TopRight, 1, matrix)[0],
		ReadUpToXInDirection(Left.x+x, Left.y+y, Left, 1, matrix)[0],
		matrix[x][y],
		ReadUpToXInDirection(Right.x+x, Right.y+y, Right, 1, matrix)[0],
		ReadUpToXInDirection(BottomLeft.x+x, BottomLeft.y+y, BottomLeft, 1, matrix)[0],
		ReadUpToXInDirection(Bottom.x+x, Bottom.y+y, Bottom, 1, matrix)[0],
		ReadUpToXInDirection(BottomRight.x+x, BottomRight.y+y, BottomRight, 1, matrix)[0],
	}
}

var (
	masUpRight   = regexp.MustCompile(`M.M.A.S.S`)
	masUpLeft    = regexp.MustCompile(`S.M.A.S.M`)
	masDownRight = regexp.MustCompile(`S.S.A.M.M`)
	masDownLeft  = regexp.MustCompile(`M.S.A.M.S`)

	allRegexes = []*regexp.Regexp{masUpRight, masUpLeft, masDownRight, masDownLeft}
)

func FindPattern(matrix [][]string) int {
	count := 0

	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if matrix[i][j] == "A" {
				snipped := Matrix1Snippet(matrix, i, j)
				snippedString := strings.Join(snipped, "")
				for _, re := range allRegexes {
					if re.MatchString(snippedString) {
						log.Println(i, j, snippedString)
						count++
					}
				}
			}
		}
	}

	return count
}
