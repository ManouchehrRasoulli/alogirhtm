package day9

import "testing"

func TestMaxRectangleSample1(t *testing.T) {
	t.Log(MaxRectangle(sample1)) // 50
}

func TestMaxRectanglePuzzle(t *testing.T) {
	t.Log(MaxRectangle(puzzle)) // 4786902990
}
