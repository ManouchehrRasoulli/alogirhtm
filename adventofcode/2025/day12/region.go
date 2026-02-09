package day12

import (
	"context"
	"sort"
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
	sort.Slice(r.shapeItems, func(i, j int) bool {
		return shapeWeight(r.shapeItems[i]) > shapeWeight(r.shapeItems[j])
	})

	return r
}

func (r region) fitAllShapes(step int, ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return false
	default:
		{
			if step == len(r.shapeItems) {
				return true
			}

			for _, shp := range r.shapeItems[step].indices {
				for y := 0; y < len(r.grid); y++ {
					for x := 0; x < len(r.grid[y]); x++ {
						if !drawShape(x, y, r.grid, shp) {
							continue
						}
						if r.fitAllShapes(step+1, ctx) {
							return true
						}
						clearShape(x, y, r.grid, shp)
					}
				}
			}

			return false
		}
	}
}

func (r region) String() string {
	builder := strings.Builder{}
	builder.WriteString("region *****\n")
	for _, shp := range r.shapeItems {
		builder.WriteString(shp.String())
	}
	builder.WriteString("grid --\n")
	for _, line := range r.grid {
		builder.WriteString(string(line))
		builder.WriteString("\n")
	}
	builder.WriteString("*****\n")
	return builder.String()
}
