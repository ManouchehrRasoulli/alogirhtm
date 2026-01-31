package day10

import (
	"fmt"
	"math"
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

func (m *MachineJoltage) PushButton(btn button) bool {
	for _, index := range btn {
		if m.currentJoltage[index]+1 > m.joltage[index] {
			fmt.Println("pushing button -->", btn, index)
			fmt.Println(m.String())
			panic("joltage exceeded")
		}
		m.currentJoltage[index]++
	}

	return true
}

func (m *MachineJoltage) DisableButtons() {
	buttonPushJoltage := func(b button, index int) bool {
		for _, idx := range b {
			if idx == index {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(m.currentJoltage); i++ {
		if m.currentJoltage[i] == m.joltage[i] {
			for bi, b := range m.buttons {
				if buttonPushJoltage(b, i) {
					m.isButtonActive[bi] = false
				}
			}
		}
	}
}

func (m *MachineJoltage) GetMinJoltageIndex() int {
	var (
		minValue = math.MaxInt
		minIndex int
	)

	for i := 0; i < len(m.joltage); i++ {
		v := m.joltage[i] - m.currentJoltage[i]
		if v <= 0 {
			continue
		}

		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	return minIndex
}

func (m *MachineJoltage) GetJoltageButtons(joltageIndex int) []button {
	var (
		buttons           []button
		buttonPushJoltage = func(b button, index int) bool {
			for _, idx := range b {
				if idx == index {
					return true
				}
			}
			return false
		}
	)

	for i := 0; i < len(m.buttons); i++ {
		if buttonPushJoltage(m.buttons[i], joltageIndex) &&
			m.isButtonActive[i] {

			buttons = append(buttons, m.buttons[i])
		}
	}

	return buttons
}

func (m *MachineJoltage) Done() bool {
	for i := 0; i < len(m.joltage); i++ {
		if m.currentJoltage[i] != m.joltage[i] {
			return false
		}
	}

	return true
}
