package day8

import "testing"

func TestConnectedCircuitsCapacityCountSample1(t *testing.T) {
	t.Log(ConnectedCircuitsCapacityCount(sample1, 10, 3)) // 40
}

func TestConnectedCircuitsCapacityCountPuzzle(t *testing.T) {
	t.Log(ConnectedCircuitsCapacityCount(puzzle, 1000, 3)) // 47040
}
