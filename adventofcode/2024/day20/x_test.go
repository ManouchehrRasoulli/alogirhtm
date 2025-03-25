package day20

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	puzzle := ReadPuzzle(data)
	t.Log(puzzle)

	t.Log(puzzle.BfsWithPath())

	t.Log(puzzle)

	cheats := puzzle.CalculateCheats()
	total := 0
	for saving, count := range cheats {
		if saving >= 8 { // at-least saved us 8 picoseconds
			total += count
		}
	}

	t.Log(cheats)
	t.Log(total)
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	puzzle := ReadPuzzle(data)
	t.Log(puzzle)

	t.Log(puzzle.BfsWithPath())

	t.Log(puzzle)

	cheats := puzzle.CalculateCheats()
	total := 0
	for saving, count := range cheats {
		if saving >= 100 { // at-least saved us 100 picoseconds
			total += count
		}
	}

	t.Log(cheats)
	t.Log(total)
}
