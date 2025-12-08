package day1

import (
	"fmt"
	"log"
	"strings"
)

func ZeroPassCount(input string, rotary *Rotary) int {
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

		for i := 0; i < steps; i++ {
			if rotary.Move(direction, 1) == 0 {
				zeros++
			}
		}
	}

	return zeros
}
