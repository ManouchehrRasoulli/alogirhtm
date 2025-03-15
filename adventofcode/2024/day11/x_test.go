package day11

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestModifyNumberPart1(t *testing.T) {
	t.Log(ModifyNumberPart1(125))    // 253000
	t.Log(ModifyNumberPart1(253000)) // 253, 0
	t.Log(ModifyNumberPart1(120501)) // 120, 501
	t.Log(ModifyNumberPart1(1))      // 2024
}

func TestStonesTestPart1(t *testing.T) {
	data, err := helper.ReadAll("input_test.part1")
	if err != nil {
		t.Fatal(err)
	}

	stones := ParseStone(data)
	t.Log(stones.MultipleBlinksPart1(25)) // 55312
}

func TestStonesPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input_puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	stones := ParseStone(data)
	t.Log(stones.MultipleBlinksPart1(25)) // 216042
}
