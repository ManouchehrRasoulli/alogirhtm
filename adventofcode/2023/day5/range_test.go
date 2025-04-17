package day5

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	input, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	seeds, mappings := ExtractMappings(input)
	t.Log(Part1(seeds, mappings)) // 35
}

func TestInputPuzzlePart1(t *testing.T) {
	input, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	seeds, mappings := ExtractMappings(input)
	t.Log(Part1(seeds, mappings)) // 111627841
}
