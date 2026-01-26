package day10

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var machineRegex = regexp.MustCompile(
	`^\[(?P<indicators>[.#]+)]\s+(?P<buttons>(?:\(\d+(?:,\d+)*\)\s*)+)\{(?P<joltage>\d+(?:,\d+)*)}$`,
)

type (
	light   rune
	button  []int
	joltage []int
)

const (
	off light = '.'
	on  light = '#'
)

type Machine struct {
	indicators []light
	buttons    []button
	joltage    joltage
	current    []light
}

func ParseMachine(s string) (*Machine, error) {
	// sample machine input
	// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
	matches := machineRegex.FindStringSubmatch(s)
	if matches == nil {
		return nil, fmt.Errorf("invalid machine indicator")
	}

	get := func(name string) string {
		for i, n := range machineRegex.SubexpNames() {
			if n == name {
				return matches[i]
			}
		}
		return ""
	}

	// Indicators
	indicatorsRaw := get("indicators")
	indicators := make([]light, len(indicatorsRaw))
	current := make([]light, len(indicatorsRaw))
	for i, c := range indicatorsRaw {
		indicators[i] = light(c)
		current[i] = off
	}

	// Buttons
	buttonsRaw := get("buttons")
	buttonMatches := regexp.MustCompile(`\((\d+(?:,\d+)*)\)`).FindAllStringSubmatch(buttonsRaw, -1)

	var buttons []button
	for _, m := range buttonMatches {
		var b button
		for _, v := range strings.Split(m[1], ",") {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			b = append(b, n)
		}
		buttons = append(buttons, b)
	}

	// Joltage
	joltageRaw := get("joltage")
	var jtg []int
	for _, v := range strings.Split(joltageRaw, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		jtg = append(jtg, n)
	}

	return &Machine{
		indicators: indicators,
		buttons:    buttons,
		joltage:    jtg,
		current:    current,
	}, nil
}

func (m *Machine) String() string {
	var sb strings.Builder

	// Indicators
	sb.WriteByte('[')
	sb.WriteString(string(m.indicators))
	sb.WriteByte(']')
	sb.WriteByte(' ')

	// Buttons
	for _, b := range m.buttons {
		sb.WriteByte('(')
		for i, v := range b {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(strconv.Itoa(v))
		}
		sb.WriteByte(')')
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

	sb.WriteString(" [")
	sb.WriteString(string(m.current))
	sb.WriteByte(']')

	return sb.String()
}

func (m *Machine) Reset() {
	for i := range m.current {
		m.current[i] = off
	}
}

func (m *Machine) PushButton(i int) {
	if i >= len(m.buttons) {
		return // couldn't push that button
	}
	btn := m.buttons[i]
	toggle := func(index int) {
		if index >= len(m.current) {
			return
		}
		ct := m.current[index]
		if ct == off {
			m.current[index] = on
		} else {
			m.current[index] = off
		}
	}

	for _, index := range btn {
		toggle(index)
	}
}

func combinations(nums []int) [][]int {
	// generate all combinations of pushing buttons
	var (
		res  [][]int
		curr []int
	)

	var backtrack func(start int)
	backtrack = func(start int) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)
	return res
}
