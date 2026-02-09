package day12

import (
	"fmt"
	"strings"
)

const (
	filled rune = '#'
	empty  rune = '.'
)

type shape struct {
	indices [][3][3]rune
}

func newShape(indic [3][3]rune) shape {
	return shape{
		indices: [][3][3]rune{
			indic,
			rotate90(indic),
			rotate180(indic),
			rotate270(indic),
		},
	}
}

func (s shape) String() string {
	builder := strings.Builder{}

	builder.WriteString("-----\n")
	for i, indic := range s.indices {
		builder.WriteString(fmt.Sprintf("-- %d\n", i))
		for _, l := range indic {
			builder.WriteString(fmt.Sprintf("%c\n", l))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
