package ady11

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
)

func Part2(universe [][]rune, rowExpand, colExpand []int) int {
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
		for j := i; j < len(galaxies); j++ {
			total += manhattanDistanceWithExpand(galaxies[i], galaxies[j], rowExpand, colExpand)
		}
	}

	return total
}

func manhattanDistanceWithExpand(a, b *helper.Location, rowExpand, colExpand []int) int {
	sum := func(a, b int, list []int) int {
		if a > b {
			a, b = b, a
		}

		v := 0
		for i := a; i < b; i++ {
			v += list[i]
		}
		return v
	}
	ax, ay := a.Get()
	bx, by := b.Get()

	dx := sum(ax, bx, rowExpand)
	dy := sum(ay, by, colExpand)

	return dx + dy
}
