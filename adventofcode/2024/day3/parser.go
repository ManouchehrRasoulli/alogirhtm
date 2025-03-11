package day3

import "regexp"

var (
	mulRgx = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
)

func FindMultiplyIndexes(input string) [][]int {
	return mulRgx.FindAllStringIndex(input, -1)
}
