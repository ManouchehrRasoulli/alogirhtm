package day2

import (
	"testing"
)

func TestReader(t *testing.T) {
	reader, err := Read("input_test.txt")
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
	reader, err := Read("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ValidCodes(reader))
}

func TestValidCodesPuzzleCodes(t *testing.T) {
	reader, err := Read("input_puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ValidCodes(reader))
}
