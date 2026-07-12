package day14

import (
	"testing"

	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
)

func TestInput0Part2(t *testing.T) {
	data, err := helper.ReadAll("input.test0.txt")
	if err != nil {
		t.Fatal(err)
	}

	ground, err := ScanGround(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("--------------- Part 2 ---------------")
	t.Log("\n", ground)

	t.Log(Part2(ground))

	t.Log("\n", ground)
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	ground, err := ScanGround(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("--------------- Part 2 ---------------")
	t.Log(Part2(ground))
}
