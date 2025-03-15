package day11

import (
	"strconv"
	"strings"
)

func ParseStone(data string) *Stones {
	s := Stones{
		numbers:      make([]int, 0),
		numbersCount: make(map[int]int),
	}

	for _, line := range strings.Split(data, "\n") {
		numbersStr := strings.Split(line, " ")
		for _, numberStr := range numbersStr {
			numberInt, _ := strconv.Atoi(numberStr)
			s.numbers = append(s.numbers, numberInt)
		}
	}

	return &s
}
