package day10

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	field, start := ReadMap(data)
	t.Log(Part1(field, start)) // 4
}

func TestInputTest2Part1(t *testing.T) {
	data, err := helper.ReadAll("input.test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	field, start := ReadMap(data)
	t.Log(Part1(field, start)) // 8
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	field, start := ReadMap(data)
	t.Log(Part1(field, start)) // 7145
}

func TestInputTest1Part2(t *testing.T) {
	data, err := helper.ReadAll("input.test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	field, start := ReadMap(data)
	printField(field)
	t.Log(Part2(field, start))
	printField(field)
}

func TestInputTest2Part2(t *testing.T) {
	data, err := helper.ReadAll("input.test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	field, start := ReadMap(data)
	printField(field)
	t.Log(Part2(field, start))
	printField(field)
}
