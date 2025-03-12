package day4

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
	"testing"
)

func TestDirectionToChar(t *testing.T) {
	t.Log(DirectionToChar[Top])
}

func TestParseInput(t *testing.T) {
	r, err := helper.Read("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(r)
	t.Log(input)
}

func TestReadUpToXInDirection(t *testing.T) {
	r, err := helper.Read("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(r)
	characters := ReadUpToXInDirection(4, 4, TopLeft, 4, input)
	t.Log(strings.Join(characters, "")) // AMXS

	characters = ReadUpToXInDirection(4, 4, TopLeft, 10, input)
	t.Log(strings.Join(characters, "")) // AMXSM

	characters = ReadUpToXInDirection(4, 4, BottomRight, 10, input)
	t.Log(strings.Join(characters, "")) // AXSAMX

	characters = ReadUpToXInDirection(0, 0, BottomRight, 10, input)
	t.Log(strings.Join(characters, "")) // MSXMAXSAMX
}

func TestNeighbors(t *testing.T) {
	r, err := helper.Read("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(r)
	ngs := Neighbors(input, 0, 0, 4)
	for _, ng := range ngs {
		t.Log(ng)
	}

	ngs = Neighbors(input, 5, 5, 4)
	for _, ng := range ngs {
		t.Log(ng)
	}
}

func TestFindPatternTestInputPart1(t *testing.T) {
	r, err := helper.Read("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(r)
	t.Log(FindPatternConstant(input)) // 18
}

func TestFindPatternPuzzleInputPart1(t *testing.T) {
	r, err := helper.Read("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(r)
	t.Log(FindPatternConstant(input)) // 2370
}

func TestMatrix1Snippet(t *testing.T) {
	r, err := helper.Read("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(r)
	data := strings.Join(Matrix1Snippet(input, 1, 2), "")
	t.Log(data) // MMSSAMMXS
	t.Log(masDownLeft.Match([]byte(data)))
}

func TestFindPatternTestInputPart2(t *testing.T) {
	r, err := helper.Read("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(r)
	t.Log(FindPattern(input)) // 9
}

func TestFindPatternPuzzleInputPart2(t *testing.T) {
	r, err := helper.Read("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := ParseInput(r)
	t.Log(FindPattern(input)) // 1908
}
