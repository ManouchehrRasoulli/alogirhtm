package day4

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"log"
	"regexp"
	"strings"
)

func ReadUpToXInDirection(x, y int, direction helper.Direction, n int, matrix [][]string) []string {
	if n == 0 {
		return []string{}
	}

	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[x]) {
		return []string{}
	}

	return append([]string{matrix[x][y]},
		ReadUpToXInDirection(x+direction.X(), y+direction.Y(), direction, n-1, matrix)...)
}

const (
	LookingFor       = "XMAS"
	RevereLookingFor = "SAMX"
)

func Neighbors(matrix [][]string, x, y int, n int) [][]string {
	ngs := make([][]string, 0)

	for _, direction := range helper.AllDirections {
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
		ReadUpToXInDirection(helper.TopLeft.X()+x, helper.TopLeft.Y()+y, helper.TopLeft, 1, matrix)[0],
		ReadUpToXInDirection(helper.Top.X()+x, helper.Top.Y()+y, helper.Top, 1, matrix)[0],
		ReadUpToXInDirection(helper.TopRight.X()+x, helper.TopRight.Y()+y, helper.TopRight, 1, matrix)[0],
		ReadUpToXInDirection(helper.Left.X()+x, helper.Left.Y()+y, helper.Left, 1, matrix)[0],
		matrix[x][y],
		ReadUpToXInDirection(helper.Right.X()+x, helper.Right.Y()+y, helper.Right, 1, matrix)[0],
		ReadUpToXInDirection(helper.BottomLeft.X()+x, helper.BottomLeft.Y()+y, helper.BottomLeft, 1, matrix)[0],
		ReadUpToXInDirection(helper.Bottom.X()+x, helper.Bottom.Y()+y, helper.Bottom, 1, matrix)[0],
		ReadUpToXInDirection(helper.BottomRight.X()+x, helper.BottomRight.Y()+y, helper.BottomRight, 1, matrix)[0],
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
