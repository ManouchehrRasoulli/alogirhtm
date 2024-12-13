package _593_find_score

import "testing"

func Test_findScore(t *testing.T) {
	v := []int{2, 1, 3, 4, 5, 2}
	t.Log(findScore(v)) // 7

	v = []int{2, 3, 5, 1, 3, 2}
	t.Log(findScore(v)) // 5
}
