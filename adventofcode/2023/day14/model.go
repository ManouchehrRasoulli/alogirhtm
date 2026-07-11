package day14

import (
	"bytes"
	"strings"
)

type value rune

const (
	empty      value = '.'
	round      value = 'O'
	recangular value = '#'
)

type ground [][]value

func (f ground) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n")
	for _, row := range f {
		buffer.WriteString(string(row))
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func TillNorth(ground ground) bool {
	changed := false

	for i := 1; i < len(ground); i++ {
		for j := 0; j < len(ground[i]); j++ {
			if ground[i][j] == round && ground[i-1][j] == empty {
				ground[i-1][j] = round
				ground[i][j] = empty
				changed = true
			}
		}
	}

	return changed
}

func CalculateWeight(ground ground) int {
	var (
		rows   = len(ground)
		weight = 0
	)

	for i := 0; i < len(ground); i++ {
		for j := 0; j < len(ground[i]); j++ {
			if ground[i][j] == round {
				weight += rows - i
			}
		}
	}

	return weight
}

func ScanGround(input string) (ground, error) {
	var (
		ground       ground
		currentField []value
	)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			ground = append(ground, currentField)
			currentField = []value{}
			continue
		}

		currentField = []value(line)
		ground = append(ground, currentField)
	}

	return ground, nil
}
