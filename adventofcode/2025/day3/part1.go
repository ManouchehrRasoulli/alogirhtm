package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func NaiveMaxJoltagePart1(input string) int {
	var (
		maxValueJoltage = func(bank string) int {
			maxVal := uint64(0)
			for i := 0; i < len(bank); i++ {
				for j := i + 1; j < len(bank); j++ {
					valStr := fmt.Sprintf("%s%s", string(bank[i]), string(bank[j]))
					val, _ := strconv.ParseUint(valStr, 10, 64)
					maxVal = max(maxVal, val)
				}
			}

			return int(maxVal)
		}

		lines = strings.Split(input, "\n")
		sum   = 0
	)

	for _, line := range lines {
		sum += maxValueJoltage(line)
	}

	return sum
}
