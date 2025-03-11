package day2

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestReader(t *testing.T) {
	reader, err := helper.Read("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	for {
		line, err := reader()
		if err != nil {
			t.Log(err)
			break
		}

		t.Log(line)
	}
}

func TestValidCodesTestCodes(t *testing.T) {
	reader, err := helper.Read("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ValidCodes(reader))
}

func TestValidCodesPuzzleCodes(t *testing.T) {
	reader, err := helper.Read("input_puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ValidCodes(reader)) // 379
}

func TestValidCodesWithToleranceTestCodes(t *testing.T) {
	reader, err := helper.Read("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ValidCodesWithTolerance(reader)) // 4
}

func TestValidCodesWithTolerancePuzzleCodes(t *testing.T) {
	reader, err := helper.Read("input_puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ValidCodesWithTolerance(reader)) // 430
}
