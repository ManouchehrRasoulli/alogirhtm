package day13

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTest0Part2(t *testing.T) {
	data, err := helper.ReadAll("input.test0.txt")
	if err != nil {
		t.Fatal(err)
	}

	fields, err := ScanFields(data)
	if err != nil {
		t.Fatal(err)
	}

	total := 0
	t.Log("--------------- Part 2 ---------------")

	for _, field := range fields {
		t.Log(field)
		val := Part2(field)
		t.Log(val)
		if val == 0 {
			t.Fatal(field.String())
		}
		total += val
	}

	t.Log("total value is:", total)
}

func TestInputPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	fields, err := ScanFields(data)
	if err != nil {
		t.Fatal(err)
	}

	total := 0
	for _, field := range fields {
		val := Part2(field)
		if val == 0 {
			t.Fatal(field.String())
		}
		total += val
	}

	t.Log("total value is:", total)
}
