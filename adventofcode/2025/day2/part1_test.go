package day2

import "testing"

func TestInvalidIdsSumSample1(t *testing.T) {
	t.Log("Testing with sample input 1")
	sum := InvalidIdsSum(sample1)
	t.Log(sum) // 1227775554
}

func TestInvalidIdsSumPuzzle(t *testing.T) {
	t.Log("Testing with input puzzle")
	sum := InvalidIdsSum(puzzle)
	t.Log(sum) // 24747430309
}
