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
	for saving, cheatCount := range cheats {
		if saving >= 8 { // at-least saved us 8 picoseconds
			total += cheatCount
		}
	}

	t.Log(cheats)
	t.Log(total) // 14
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
	for saving, cheatCount := range cheats {
		if saving >= 100 { // at-least saved us 100 picoseconds
			total += cheatCount
		}
	}

	t.Log(cheats)
	t.Log(total) // 1521
}

func TestInputTest1Part2(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	puzzle := ReadPuzzle(data)
	t.Log(puzzle)

	t.Log(puzzle.BfsWithPath())

	t.Log(puzzle)

	cheats := puzzle.CalculateCheatsWith20Picoseconds()
	total := 0
	for saving, cheatCount := range cheats {
		if saving >= 50 { // at-least saved us 50 picoseconds
			total += cheatCount
		}
	}

	t.Log(cheats)
	t.Log(total) // 285
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	puzzle := ReadPuzzle(data)
	t.Log(puzzle)

	t.Log(puzzle.BfsWithPath())

	t.Log(puzzle)

	cheats := puzzle.CalculateCheatsWith20Picoseconds()
	total := 0
	for saving, cheatCount := range cheats {
		if saving >= 100 { // at-least saved us 100 picoseconds
			total += cheatCount
		}
	}

	t.Log(cheats)
	t.Log(total) // 1013106
}
