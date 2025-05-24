package day8

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	fileInput, err := helper.ReadAll("input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(fileInput)
	t.Log(input.Part1()) // 2
}

func TestInputTestPart2(t *testing.T) {
	fileInput, err := helper.ReadAll("input.test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(fileInput)
	t.Log(input.Part1()) // 6
}

func TestInputPuzzlePart1(t *testing.T) {
	fileInput, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(fileInput)
	t.Log(input.Part1()) // 19951
}
