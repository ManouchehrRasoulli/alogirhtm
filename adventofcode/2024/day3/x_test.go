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
	multiplies, err := ReadMultiplies(indexes, program, false)
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
	multiplies, err := ReadMultiplies(indexes, program, false)
	if err != nil {
		t.Fatal(err)
	}

	sum := 0
	for _, i := range multiplies {
		sum += i.Value()
	}

	t.Log(sum) // 164730528
}

func TestMultiplyTestInputDoDont(t *testing.T) {
	program, err := helper.ReadAll("input_test_part2.txt")
	if err != nil {
		t.Fatal(err)
	}

	multiplyIndexes := FindMultiplyIndexes(program)
	multiplies, err := ReadMultiplies(multiplyIndexes, program, true)
	if err != nil {
		t.Fatal(err)
	}

	sum := 0
	for _, i := range multiplies {
		sum += i.Value()
	}

	t.Log(sum) // 48
}

func TestMultiplyPuzzleInputDoDont(t *testing.T) {
	program, err := helper.ReadAll("input_puzzle_part2.txt")
	if err != nil {
		t.Fatal(err)
	}

	multiplyIndexes := FindMultiplyIndexes(program)
	multiplies, err := ReadMultiplies(multiplyIndexes, program, true)
	if err != nil {
		t.Fatal(err)
	}

	sum := 0
	for _, i := range multiplies {
		sum += i.Value()
	}

	t.Log(sum) // 70478672
}
