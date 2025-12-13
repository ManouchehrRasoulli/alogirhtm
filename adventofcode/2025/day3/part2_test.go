package day3

import "testing"

func TestMaxJoltagePart2Sample(t *testing.T) {
	sum := MaxJoltagePart2(sample1)
	t.Log(sum) // 3121910778619
}

func TestMaxJoltagePart2Puzzle(t *testing.T) {
	sum := MaxJoltagePart2(puzzle)
	t.Log(sum) // 166345822896410
}
