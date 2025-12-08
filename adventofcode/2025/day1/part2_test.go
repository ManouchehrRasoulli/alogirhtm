package day1

import "testing"

func TestZeroPassCountSample1(t *testing.T) {
	rotary := NewRotary(50)
	zeroCount := ZeroPassCount(sample1, rotary)
	if zeroCount != 6 {
		t.Fatal("expected 6, got ", zeroCount)
	}
	t.Log(zeroCount)
}

func TestZeroPassCountPuzzle(t *testing.T) {
	rotary := NewRotary(50)
	t.Log(ZeroPassCount(puzzle, rotary))
}
