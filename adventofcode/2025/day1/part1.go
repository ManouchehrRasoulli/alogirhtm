package day1

import (
	"fmt"
	"log"
	"strings"
)

func ZeroCount(input string, rotary *Rotary) int {
	var (
		lines = strings.Split(input, "\n")
		zeros = 0
	)

	for _, line := range lines {
		var (
			direction Direction
			steps     int
		)

		_, err := fmt.Sscanf(line, "%1s%d", &direction, &steps)
		if err != nil {
			log.Fatal(err, " input line is :: ", line)
		}

		if rotary.Move(direction, steps) == 0 {
			zeros++
		}
	}

	return zeros
}
