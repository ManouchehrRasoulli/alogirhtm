package day9

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	items := ReadData(data)
	t.Log(items)

	t.Log(Part1(items)) // 114
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	items := ReadData(data)
	t.Log(items)

	t.Log(Part1(items)) // 1684566095
}

func TestInputTest1Part2(t *testing.T) {
	data, err := helper.ReadAll("input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	items := ReadData(data)
	t.Log(items)

	t.Log(Part2(items)) // 2
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	items := ReadData(data)
	t.Log(items)

	t.Log(Part2(items)) // 1136
}
