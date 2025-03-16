package day12

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestParseAndFloodTestPart1(t *testing.T) {
	input, err := helper.ReadAll("input_test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	garden := ParseGarden(input)
	t.Log(garden)

	t.Log(garden.Flood(Location{
		x: 0,
		y: 0,
	})) // 4, 10

	t.Log(garden.Flood(Location{
		x: 0,
		y: 3,
	})) // 0, 0

	t.Log(garden.Flood(Location{
		x: 1,
		y: 3,
	})) // 1, 4

	t.Log(garden.Flood(Location{
		x: 1,
		y: 2,
	})) // 4, 10

	t.Log(garden.Flood(Location{
		x: 2,
		y: 1,
	})) // 4, 8

	t.Log(garden.Flood(Location{
		x: 3,
		y: 1,
	})) // 4, 8
}

func TestFencingTest2Part1(t *testing.T) {
	input, err := helper.ReadAll("input_test2.part1")
	if err != nil {
		t.Fatal(err)
	}

	garden := ParseGarden(input)
	t.Log(garden)

	t.Log(garden.FencesAmountPart1()) // 772
}

func TestFencingTest3Part1(t *testing.T) {
	input, err := helper.ReadAll("input_test3.part1")
	if err != nil {
		t.Fatal(err)
	}

	garden := ParseGarden(input)
	t.Log(garden)

	t.Log(garden.FencesAmountPart1()) // 1930
}

func TestFencingPuzzle3Part1(t *testing.T) {
	input, err := helper.ReadAll("input_puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	garden := ParseGarden(input)
	t.Log(garden)

	t.Log(garden.FencesAmountPart1()) // 1465968
}
