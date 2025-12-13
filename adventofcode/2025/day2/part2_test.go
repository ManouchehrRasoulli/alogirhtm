package day2

import "testing"

func TestInvalidIdsRepeatedSumSample1(t *testing.T) {
	t.Log("Testing with sample input 1")
	sum := InvalidIdsRepeatedSum(sample1)
	t.Log(sum) // 4174379265
}

func TestInvalidIdsRepeatedSumPuzzle(t *testing.T) {
	t.Log("Testing with input puzzle")
	sum := InvalidIdsRepeatedSum(puzzle)
	t.Log(sum) // 30962646823
}
