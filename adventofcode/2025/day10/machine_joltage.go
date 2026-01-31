package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type MachineJoltage struct {
	buttons        []button
	isButtonActive []bool
	joltage        joltage
	currentJoltage joltage
}

func ParseMachineJoltage(s string) (*MachineJoltage, error) {
	m, err := ParseMachine(s)
	if err != nil {
		return nil, err
	}

	isButtonActive := make([]bool, len(m.buttons))
	for i := 0; i < len(m.buttons); i++ {
		isButtonActive[i] = true
	}

	return &MachineJoltage{
		buttons:        m.buttons,
		isButtonActive: isButtonActive,
		joltage:        m.joltage,
		currentJoltage: make(joltage, len(m.joltage)),
	}, nil
}

func (m *MachineJoltage) String() string {
	var sb strings.Builder

	// Buttons
	for bi, b := range m.buttons {
		sb.WriteByte('(')
		for i, v := range b {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(strconv.Itoa(v))
		}
		sb.WriteByte(')')
		sb.WriteByte('-')
		sb.WriteString(fmt.Sprintf("%t", m.isButtonActive[bi]))
		sb.WriteByte(' ')
	}

	// Joltage
	sb.WriteByte('{')
	for i, v := range m.joltage {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(v))
	}
	sb.WriteByte('}')

	sb.WriteString(" {")
	for i, v := range m.currentJoltage {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(v))
	}
	sb.WriteByte('}')

	return sb.String()
}

func (m *MachineJoltage) Done() bool {
	for i := 0; i < len(m.joltage); i++ {
		if m.currentJoltage[i] != m.joltage[i] {
			return false
		}
	}
	return true
}
