package helper

import (
	"log"
	"strconv"
)

func ArrayStrToArrayInt(nums []string) []int {
	var (
		output = make([]int, 0)
	)

	for _, n := range nums {
		if n == "" {
			continue
		}

		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err, nums, n)
		}

		output = append(output, num)
	}

	return output
}
