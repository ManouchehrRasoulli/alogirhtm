package day1

import (
	_ "embed"
	"testing"
)

//go:embed sample1.txt
var sample1 string

//go:embed puzzle.txt
var puzzle string

func TestZeroCountSample1(t *testing.T) {
	rotary := NewRotary(50)
	zeroCount := ZeroCount(sample1, rotary)
	if zeroCount != 3 {
		t.Fatal("expected 3, got ", zeroCount)
	}
	t.Log(zeroCount)
}

func TestZeroCountPuzzle(t *testing.T) {
	rotary := NewRotary(50)
	t.Log(ZeroCount(puzzle, rotary))
}
