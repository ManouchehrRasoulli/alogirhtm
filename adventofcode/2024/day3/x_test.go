package day3

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestMultiplyTestInput(t *testing.T) {
	program, err := helper.ReadAll("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	indexes := FindMultiplyIndexes(program)
	multiplies, err := ReadMultiplies(indexes, program)
	if err != nil {
		t.Fatal(err)
	}

	sum := 0
	for _, i := range multiplies {
		sum += i.Value()
	}

	t.Log(sum) // 161
}

func TestMultiplyPuzzleInput(t *testing.T) {
	program, err := helper.ReadAll("input_puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	indexes := FindMultiplyIndexes(program)
	multiplies, err := ReadMultiplies(indexes, program)
	if err != nil {
		t.Fatal(err)
	}

	sum := 0
	for _, i := range multiplies {
		sum += i.Value()
	}

	t.Log(sum) // 164730528
}
