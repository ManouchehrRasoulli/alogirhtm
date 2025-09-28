package day12

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTest0Part2(t *testing.T) {
	data, err := helper.ReadAll("input.test0.txt")
	if err != nil {
		t.Fatal(err)
	}

	items := readInput(data)
	t.Log(Part2(items)) // 525152
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	items := readInput(data)
	t.Log(Part2(items)) // 18093821750095
}
