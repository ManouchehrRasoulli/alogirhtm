package day10

import (
	"fmt"
	"strings"
)

func solveMachineJoltage(m *MachineJoltage) int {
	var (
		count         int
		maxButtonFunc = func(buttons []button) int {
			var (
				bt      button
				btIndex int
			)

			for i, b := range buttons {
				if len(b) > len(bt) {
					bt = b
					btIndex = i
				}
			}

			return btIndex
		}
	)

	for !m.Done() {
		minJoltageIndex := m.GetMinJoltageIndex()
		buttons := m.GetJoltageButtons(minJoltageIndex)
		if len(buttons) == 0 {
			fmt.Println(m)
			panic("no buttons for pushing")
		}
		maxButtonIndex := maxButtonFunc(buttons)
		if !m.PushButton(buttons[maxButtonIndex]) {
			fmt.Println(m)
			panic("couldn't push button")
		}
		m.DisableButtons()
		count++
	}

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
