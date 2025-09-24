package ady11

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
)

func Part1(universe [][]rune) int {
	var (
		total    = 0
		galaxies = make([]*helper.Location, 0)
	)

	for i, row := range universe {
		for j, col := range row {
			if col == galaxy {
				galaxies = append(galaxies, helper.NewLocation(i, j))
			}
		}
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += manhattanDistance(galaxies[i], galaxies[j])
		}
	}

	return total
}

func manhattanDistance(a, b *helper.Location) int {
	ax, ay := a.Get()
	bx, by := b.Get()

	dx := ax - bx
	if dx < 0 {
		dx = -dx
	}
	dy := ay - by
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}
