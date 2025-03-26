package day21

import (
	"math"
	"regexp"
	"strconv"
)

var (
	numberRegex = regexp.MustCompile(`[0-9]+`)
)

func ComputeValues(sequence string, finalPath [][]rune) int {
	numStr := numberRegex.FindAllString(sequence, 1)[0]
	num, _ := strconv.Atoi(numStr)

	minLen := math.MaxInt
	for _, v := range finalPath {
		if len(v) < minLen {
			minLen = len(v)
		}
	}

	return num * minLen
}
