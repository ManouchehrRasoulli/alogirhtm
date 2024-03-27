package missing

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	ls := []int{1, 2, 0}
	t.Log(FistMissingPositive(ls))
}
