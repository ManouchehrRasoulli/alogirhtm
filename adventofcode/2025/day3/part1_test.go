package day3

import "testing"

func TestMaxJoltagePart1Sample(t *testing.T) {
	sum := MaxJoltagePart1(sample1)
	t.Log(sum) // 357
}

func TestMaxJoltagePart1Puzzle(t *testing.T) {
	sum := MaxJoltagePart1(puzzle)
	t.Log(sum) // 16812
}
