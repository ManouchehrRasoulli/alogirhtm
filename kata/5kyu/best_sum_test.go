package _kyu

import "testing"

func TestChooseBestSum(t *testing.T) {
	ts := []int{91, 74, 73, 85, 73, 81, 87}
	t.Log(ChooseBestSum(230, 3, ts)) // 228
}
