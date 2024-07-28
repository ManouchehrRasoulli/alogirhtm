package _kyu

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	t.Log(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	t.Log(MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}))
}
