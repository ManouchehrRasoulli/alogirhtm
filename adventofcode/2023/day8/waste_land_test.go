package day8

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTest1Part1(t *testing.T) {
	fileInput, err := helper.ReadAll("input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(fileInput)
	t.Log(input.Part1()) // 2
}

func TestInputTest2Part1(t *testing.T) {
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

func TestInputTest1Part2(t *testing.T) {
	fileInput, err := helper.ReadAll("input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(fileInput)
	t.Log(input.Part2()) // 2
}

func TestInputTest2Part2(t *testing.T) {
	fileInput, err := helper.ReadAll("input.test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(fileInput)
	t.Log(input.Part2()) // 6
}

func TestInputTest3Part2(t *testing.T) {
	fileInput, err := helper.ReadAll("input.test3.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(fileInput)
	t.Log(input.Part2()) // 6
}

func TestInputPuzzlePart2(t *testing.T) {
	fileInput, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(fileInput)
	t.Log(input.Part2()) // 16342438708751
}
