package day10

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	field, err := helper.ReadAll("input_test.part1")
	if err != nil {
		t.Fatal(err)
	}

	island := ParseIsland(field)
	t.Log(island)

	t.Log(island.TrailheadScoresSumPart1()) // 36
}

func TestInputPuzzlePart1(t *testing.T) {
	field, err := helper.ReadAll("input_puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	island := ParseIsland(field)
	t.Log(island)

	t.Log(island.TrailheadScoresSumPart1()) // 659
}

func TestInputTestPart2(t *testing.T) {
	field, err := helper.ReadAll("input_test.part1")
	if err != nil {
		t.Fatal(err)
	}

	island := ParseIsland(field)
	t.Log(island)

	t.Log(island.TrailheadScoresSumPart2()) // 81
}

func TestInputPuzzlePart2(t *testing.T) {
	field, err := helper.ReadAll("input_puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	island := ParseIsland(field)
	t.Log(island)

	t.Log(island.TrailheadScoresSumPart2()) // 1463
}
