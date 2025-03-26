package day21

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
	"testing"
)

func TestInputTest1Part1(t *testing.T) {
	input, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Log(err)
	}

	final := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		keypad := NewKeypad(WithKeyPad(line))
		output := keypad.KeyToDirections()
		final += ComputeValues(line, output)
	}

	t.Log(final) // 126384
}

func TestInputPuzzlePart1(t *testing.T) {
	input, err := helper.ReadAll("input.puzzle")
	if err != nil {
		t.Log(err)
	}

	final := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		keypad := NewKeypad(WithKeyPad(line))
		output := keypad.KeyToDirections()
		final += ComputeValues(line, output)
	}

	t.Log(final) // (133834, 141162) // (actual -- 137870)
}
