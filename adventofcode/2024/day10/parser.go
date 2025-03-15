package day10

import (
	"strconv"
	"strings"
)

func ParseIsland(data string) *Island {
	is := Island{
		field:      make([][]int, 0),
		trailheads: make(map[Location]struct{}),
	}

	for i, line := range strings.Split(data, "\n") {
		items := strings.Split(line, "")
		numbers := make([]int, 0)
		for j, it := range items {
			number, _ := strconv.Atoi(it)
			numbers = append(numbers, number)
			if number == 0 {
				is.trailheads[Location{i, j}] = struct{}{}
			}
		}

		is.field = append(is.field, numbers)
	}

	return &is
}
