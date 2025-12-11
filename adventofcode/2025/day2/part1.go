package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func InvalidIdsSum(input string) int {
	isInvalidId := func(x string) bool {
		if len(x)%2 != 0 {
			return false
		}
		return x[:len(x)/2] == x[len(x)/2:]
	}
	var sum int
	rangeIds := strings.Split(input, ",")
	for _, rangeId := range rangeIds {
		parts := strings.Split(rangeId, "-")
		startStr, endStr := parts[0], parts[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)

		for i := start; i <= end; i++ {
			if isInvalidId(fmt.Sprintf("%d", i)) {
				sum += i
			}
		}
	}
	return sum
}
