package day13

import (
	"testing"

	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
)

func TestInputTest0Part1(t *testing.T) {
	data, err := helper.ReadAll("input.test0.txt")
	if err != nil {
		t.Fatal(err)
	}

	fields, err := ScanFields(data)
	if err != nil {
		t.Fatal(err)
	}

	total := 0
	t.Log("--------------- Part 1 ---------------")

	for _, field := range fields {
		t.Log(field)
		val := Part1(field)
		t.Log(val)
		if val == 0 {
			t.Fatal(field.String())
		}
		total += val
	}

	t.Log("total value is:", total)
}

func TestInputPuzzlePart1(t *testing.T) {
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
		val := Part1(field)
		if val == 0 {
			t.Fatal(field.String())
		}
		total += val
	}

	t.Log("total value is:", total)
}
