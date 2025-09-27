package day12

import (
	"fmt"
	"strconv"
	"strings"
)

type value rune

const (
	operational value = '.'
	damaged     value = '#'
	unknown     value = '?'
)

type dataLines struct {
	Data  []value
	Match []int
}

func (d dataLines) String() string {
	return fmt.Sprintf("%s %v", string(d.Data), d.Match)
}

func readInput(in string) []dataLines {
	data := make([]dataLines, 0)

	// the input sample is
	// ???.### 1,1,3

	for _, line := range strings.Split(in, "\n") {
		lineItem := dataLines{
			Data:  make([]value, 0),
			Match: make([]int, 0),
		}
		parts := strings.Split(line, " ")

		// read part 1
		for _, v := range parts[0] {
			lineItem.Data = append(lineItem.Data, value(v))
		}

		// read part 2
		for _, v := range strings.Split(parts[1], ",") {
			vi, _ := strconv.Atoi(v)
			lineItem.Match = append(lineItem.Match, vi)
		}

		data = append(data, lineItem)
	}

	return data
}
