package day3

import "regexp"

var (
	mulWithDoDontRgx = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
)

const (
	DoFlag   = "do()"
	DontFlag = "don't()"
	MulFlag  = "mul(+d,+d)"
)

func FindMultiplyIndexes(input string) [][]int {
	return mulWithDoDontRgx.FindAllStringIndex(input, -1)
}
