package day10

import (
	"fmt"
	"strings"
)

func solveMachineJoltage(m *MachineJoltage) int {
	var count int

	fmt.Println(m)

	return count
}

func Part2(input string) int {
	var (
		lines = strings.Split(input, "\n")
		total = 0
	)

	for _, line := range lines {
		m, err := ParseMachineJoltage(line)
		if err != nil {
			panic(err)
		}

		total += solveMachineJoltage(m)
	}

	return total
}
