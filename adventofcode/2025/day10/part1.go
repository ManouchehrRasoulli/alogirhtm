package day10

import (
	"math"
	"strings"
)

func solveMachine(m *Machine) int {
	var (
		btns  = make([]int, 0)
		pushs = math.MaxInt32
	)
	for i := 0; i < len(m.buttons); i++ {
		btns = append(btns, i)
	}

	cmb := combinations(btns)
	for _, cm := range cmb {
		m.Reset()
		for _, bt := range cm {
			m.PushButton(bt)
		}
		if string(m.current) == string(m.indicators) {
			pushs = min(pushs, len(cm))
		}
	}

	if pushs != math.MaxInt32 {
		return pushs
	}

	return 0
}

func Part1(input string) int {
	var (
		lines = strings.Split(input, "\n")
		total = 0
	)

	for _, line := range lines {
		m, err := ParseMachine(line)
		if err != nil {
			panic(err)
		}

		total += solveMachine(m)
	}

	return total
}
