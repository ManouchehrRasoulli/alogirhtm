package _kyu

import "testing"

func TestTwoSum(t *testing.T) {
	s := TwoSum([]int{1, 2, 3}, 4)
	t.Log(s)
	s = TwoSum([]int{1234, 5678, 9012}, 14690)
	t.Log(s)
	s = TwoSum([]int{2, 2, 3}, 4)
	t.Log(s)
	s = TwoSum([]int{1, 1, 3}, 5)
	t.Log(s)
}
