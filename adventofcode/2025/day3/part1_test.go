package day3

import "testing"

func TestNaiveMaxJoltagePart1Sample(t *testing.T) {
	sum := NaiveMaxJoltagePart1(sample1)
	t.Log(sum) // 357
}

func TestNaiveMaxJoltagePart1Puzzle(t *testing.T) {
	sum := NaiveMaxJoltagePart1(puzzle)
	t.Log(sum) // 16812
}
