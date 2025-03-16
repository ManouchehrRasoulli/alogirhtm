package day12

import "strings"

func ParseGarden(input string) *Garden {
	g := Garden{
		Plots:   make([][]rune, 0),
		Visited: make(map[Location]struct{}),
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		g.Plots = append(g.Plots, []rune(line))
	}

	return &g
}
