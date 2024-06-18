package _kyu

import (
	"strconv"
)

func numOne(s string) int {
	var oneSum int
	for _, b := range s {
		if b == '1' {
			oneSum += 1
		}
	}
	return oneSum
}

func NextHigherFirstAttempt(n int) int {
	nBinary := strconv.FormatInt(int64(n), 2)
	nOneSum := numOne(nBinary)
	for {
		n += 1
		nBinary = strconv.FormatInt(int64(n), 2)
		if numOne(nBinary) == nOneSum {
			return n
		}
	}
}

func NextHigher(n int) int {
	r1 := n & -n                   // rightmost non-zero bit
	next := n + r1                 // next with the same number of ones
	diff := ((n ^ next) >> 2) / r1 // shifted differences

	return next | diff
}
