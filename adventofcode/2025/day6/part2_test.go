package day6

import "testing"

func TestSolveHorizontallyAndSumSample1(t *testing.T) {
	t.Log(SolveHorizontallyAndSum(sample1)) // 3263827
}

func TestSolveHorizontallyAndSumPuzzle(t *testing.T) {
	t.Log(SolveHorizontallyAndSum(puzzle)) // 9581313737063
}
