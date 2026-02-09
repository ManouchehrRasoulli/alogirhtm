package day12

import (
	"fmt"
	"strings"
)

type region struct {
	shapeItems []shape
	grid       [][]rune
}

func newRegion(grid [][]rune, shapes []shape, quantities []int) region {
	r := region{
		shapeItems: make([]shape, 0),
		grid:       grid,
	}
	for i := 0; i < len(quantities); i++ {
		for j := 0; j < quantities[i]; j++ {
			r.shapeItems = append(r.shapeItems, shapes[i])
		}
	}

	return r
}

func (r region) String() string {
	builder := strings.Builder{}
	builder.WriteString("region ---\n")
	builder.WriteString(fmt.Sprintf("shapes: %v\n", r.shapeItems))
	builder.WriteString(fmt.Sprintf("grid: %v\n", r.grid))
	return builder.String()
}
