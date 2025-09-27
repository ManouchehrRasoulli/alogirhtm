package day12

import (
	"fmt"
	"slices"
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

func (d dataLines) DataMatches() bool {
	var (
		currentPattern = make([]int, 0)
		currentCount   = 0
	)

	for _, v := range d.Data {
		if v != damaged && currentCount != 0 {
			currentPattern = append(currentPattern, currentCount)
			currentCount = 0
			continue
		}

		if v == damaged {
			currentCount++
		}
	}

	if currentCount != 0 {
		currentPattern = append(currentPattern, currentCount)
	}

	return slices.Compare(currentPattern, d.Match) == 0
}

func (d dataLines) Unfold(count int) dataLines {
	newDataLine := dataLines{
		Data:  make([]value, 0),
		Match: make([]int, 0),
	}

	for i := 0; i < count; i++ {
		newDataLine.Data = append(newDataLine.Data, d.Data...)
		if i+1 != count {
			newDataLine.Data = append(newDataLine.Data, unknown)
		}
		newDataLine.Match = append(newDataLine.Match, d.Match...)
	}

	return newDataLine
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
