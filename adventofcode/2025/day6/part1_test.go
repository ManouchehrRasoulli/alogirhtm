package day6

import "testing"

func TestSolveAndSumSample1(t *testing.T) {
	t.Log(SolveAndSum(sample1)) // 4277556
}

func TestSolveAndSumPuzzle(t *testing.T) {
	t.Log(SolveAndSum(puzzle)) // 4722948564882
}
