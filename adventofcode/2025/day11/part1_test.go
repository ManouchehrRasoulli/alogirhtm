package day11

import "testing"

func TestUniquePathsSample1(t *testing.T) {
	t.Log(uniquePaths(sample1)) // 5
}

func TestUniquePathsPuzzle(t *testing.T) {
	t.Log(uniquePaths(puzzle)) // 696
}
