package ady11

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestPart1Test0(t *testing.T) {
	input, err := helper.ReadAll("./input.test0.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("original universe")
	universe := ReadUniverse(input)
	PrintUniverse(universe)

	t.Log("gravity affected --> expanded universe")
	universe = ExpandUniverse(universe)
	PrintUniverse(universe)

	t.Log(Part1(universe))
}

func TestPart1Test1(t *testing.T) {
	input, err := helper.ReadAll("./input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("original universe")
	universe := ReadUniverse(input)
	PrintUniverse(universe)

	t.Log("gravity affected --> expanded universe")
	universe = ExpandUniverse(universe)
	PrintUniverse(universe)

	t.Log(Part1(universe))
}

func TestPart1Puzzle(t *testing.T) {
	input, err := helper.ReadAll("./input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	universe := ReadUniverse(input)

	universe = ExpandUniverse(universe)

	t.Log(Part1(universe))
}

func TestPart2Test1(t *testing.T) {
	input, err := helper.ReadAll("./input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	universe := ReadUniverse(input)
	PrintUniverse(universe)

	rowExpand, colExpand := ExpandNUniverse(universe, 100)
	t.Log(rowExpand, colExpand)
	t.Log(Part2(universe, rowExpand, colExpand))
}

func TestPart2Puzzle(t *testing.T) {
	input, err := helper.ReadAll("./input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	universe := ReadUniverse(input)
	rowExpand, colExpand := ExpandNUniverse(universe, 1_000_000)
	t.Log(Part2(universe, rowExpand, colExpand))
}
