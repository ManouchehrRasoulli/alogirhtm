package helper

import "testing"

func TestRemoveAndReturnIndex(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	t.Log(RemoveAndReturnIndex(s, 2))

	s = []int{1}
	t.Log(RemoveAndReturnIndex(s, 0))
}
