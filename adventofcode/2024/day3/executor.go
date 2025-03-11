package day3

import (
	"fmt"
	"strconv"
	"strings"
)

type Multiply struct {
	A int
	B int
}

func MultiplyFromString(s string) (*Multiply, error) {
	part1 := strings.Split(s, "(")[1]
	part2 := strings.Split(part1, ")")[0]
	values := strings.Split(part2, ",")

	a, err := strconv.Atoi(values[0])
	if err != nil {
		return nil, fmt.Errorf("invalid multiply: %s", s)
	}
	b, err := strconv.Atoi(values[1])
	if err != nil {
		return nil, fmt.Errorf("invalid multiply: %s", s)
	}

	return &Multiply{A: a, B: b}, nil
}

func (m Multiply) Value() int {
	return m.A * m.B
}

func ReadMultiplies(indexes [][]int, input string, doDontEnable bool) ([]Multiply, error) {
	multiplies := make([]Multiply, 0)

	care := true

indexesLoop:
	for _, index := range indexes {
		v := input[index[0]:index[1]]

		switch v {
		case DontFlag:
			care = false
			continue indexesLoop
		case DoFlag:
			care = true
			continue indexesLoop
		}

		if care || !doDontEnable {
			m, err := MultiplyFromString(v)
			if err != nil {
				return nil, err
			}

			multiplies = append(multiplies, *m)
		}
	}
	return multiplies, nil
}
