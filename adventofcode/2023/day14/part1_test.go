package day14

import (
	"testing"

	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
)

func TestInput0Part1(t *testing.T) {
	data, err := helper.ReadAll("input.test0.txt")
	if err != nil {
		t.Fatal(err)
	}

	ground, err := ScanGround(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("--------------- Part 1 ---------------")
	t.Log("\n", ground)

	t.Log(Part1(ground))

	t.Log("\n", ground)
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	ground, err := ScanGround(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("--------------- Part 1 ---------------")
	t.Log(Part1(ground))
}
