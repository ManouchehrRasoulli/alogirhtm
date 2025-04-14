package day3

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	p := NewPuzzle(data)
	t.Log(p)

	t.Log(p.Part1()) // 4361

	t.Log(p)
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	p := NewPuzzle(data)
	t.Log(p)

	t.Log(p.Part1()) // 560670

	t.Log(p)
}
