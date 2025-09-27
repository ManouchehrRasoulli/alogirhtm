package day12

import (
	"fmt"
)

func Part1(lines []dataLines) int {
	var permutations int

	for _, line := range lines {
		fmt.Println(line)
		permutations += countPermutations(0, line)
	}

	return permutations
}

func countPermutations(position int, line dataLines) int {
	if position == len(line.Data) {
		if line.DataMatches() {
			return 1
		}
		return 0
	}

	if line.Data[position] != unknown {
		return countPermutations(position+1, line)
	}

	line.Data[position] = operational
	p1 := countPermutations(position+1, line)
	line.Data[position] = damaged
	p2 := countPermutations(position+1, line)
	line.Data[position] = unknown
	return p1 + p2
}
