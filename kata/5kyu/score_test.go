package _kyu

import "testing"

func TestScore(t *testing.T) {
	t.Log(Score([5]int{2, 3, 4, 6, 2}))
	t.Log(Score([5]int{1, 1, 1, 3, 1}))
	t.Log(Score([5]int{2, 4, 4, 5, 4}))
}
