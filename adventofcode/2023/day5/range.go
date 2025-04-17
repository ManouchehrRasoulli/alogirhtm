package day5

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"log"
	"math"
	"strings"
)

const (
	noValue = -1
)

type Range struct {
	Dest   int // first part
	Source int
	Len    int
}

func (r Range) Map(input int) int {
	// map input number to destination range number
	if input >= r.Source && input <= r.Source+r.Len {
		dif := input - r.Source
		return r.Dest + dif
	}

	return noValue // range didn't match the input
}

type Mapping struct {
	from   string
	to     string
	ranges []Range
}

func (m Mapping) Map(input int) int {
	var (
		dest int
	)

	for _, r := range m.ranges {
		dest = r.Map(input)
		if dest != noValue {
			return dest
		}
	}

	return input // didn't match any range value
}

func ExtractMappings(input string) ([]int, []Mapping) {
	var (
		seeds    = make([]int, 0)
		mappings = make([]Mapping, 0)
		lines    = strings.Split(input, "\n")
		mapping  = Mapping{}
	)

	seeds = helper.ArrayStrToArrayInt(strings.Split(strings.Split(lines[0], ":")[1], " "))

	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			mappings = append(mappings, mapping)
			mapping = Mapping{}
			continue
		}

		if strings.Contains(lines[i], "map") {
			parts := strings.Split(lines[i], "-")
			mapping.from = parts[0]
			mapping.to = parts[2]
			continue
		}

		nums := helper.ArrayStrToArrayInt(strings.Split(lines[i], " "))
		if len(nums) != 3 {
			log.Fatal("invalid range number ", lines[i], " ", nums, " ", len(nums))
		}

		numRange := Range{
			Dest:   nums[0],
			Source: nums[1],
			Len:    nums[2],
		}
		mapping.ranges = append(mapping.ranges, numRange)
	}

	return seeds, mappings
}

func Part1(seeds []int, mappings []Mapping) int {
	var (
		minValue = math.MaxInt
	)

	for _, seed := range seeds {
		prev := seed
		for _, m := range mappings {
			prev = m.Map(prev)
		}
		if prev < minValue {
			minValue = prev
		}
	}

	return minValue
}
