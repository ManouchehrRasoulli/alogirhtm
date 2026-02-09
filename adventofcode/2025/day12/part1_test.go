package day12

import "testing"

func TestFitShapesSample1(t *testing.T) {
	t.Log(FitShapes(sample1)) // 2
}

func TestFitShapesPuzzle(t *testing.T) {
	t.Log(FitShapes(puzzle))
}
