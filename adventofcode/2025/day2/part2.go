package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func InvalidIdsRepeatedSum(input string) int {
	isInvalidId := func(x string) bool {
		for k := 1; k <= len(x)/2; k++ {
			if len(x)%k != 0 {
				continue
			}

			blocks := make([]string, 0)
			for i := 0; i < len(x); i += k {
				blocks = append(blocks, x[i:i+k])
			}

			isAllSame := true
			for i := 1; i < len(blocks); i++ {
				if blocks[i] != blocks[i-1] {
					isAllSame = false
					break
				}
			}

			if isAllSame {
				return true
			}
		}

		return false
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
