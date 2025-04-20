package day6

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"regexp"
	"strings"
)

const (
	timing    = "Time"
	distances = "Distance"
)

var (
	numberRegex = regexp.MustCompile(`[0-9]+`)
)

func ReadRaceAndTiming(input string) ([]int, []int) {
	var (
		time     = make([]int, 0)
		distance = make([]int, 0)
		lines    = strings.Split(input, "\n")
	)

	for _, line := range lines {
		if strings.HasPrefix(line, timing) {
			time = helper.ArrayStrToArrayInt(numberRegex.FindAllString(line, -1))
		}

		if strings.HasPrefix(line, distances) {
			distance = helper.ArrayStrToArrayInt(numberRegex.FindAllString(line, -1))
		}
	}

	return time, distance
}

func Part1(time, distance []int) int {
	var (
		total = 1
	)

	for i, t := range time {
		var (
			di = distance[i]
			tc = 0
		)

		// compute the timing for each combination
		for wt := 0; wt <= t; wt++ {
			rt := t - wt
			sd := rt * wt
			if sd > di {
				tc++
			}
		}

		total *= tc
	}

	return total
}
