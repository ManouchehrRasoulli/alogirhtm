package day3

import (
	"strings"
)

func MaxJoltagePart1(input string) int {
	var (
		maxValueJoltage = func(bank string) int {
			if len(bank) < 2 {
				return 0
			}
			maxVal := 0
			maxRight := int(bank[len(bank)-1] - '0')
			for i := len(bank) - 2; i >= 0; i-- {
				current := int(bank[i] - '0')
				val := current*10 + maxRight
				if val > maxVal {
					maxVal = val
				}
				if current > maxRight {
					maxRight = current
				}
			}
			return maxVal
		}

		lines = strings.Split(input, "\n")
		sum   = 0
	)

	for _, line := range lines {
		sum += maxValueJoltage(line)
	}

	return sum
}
