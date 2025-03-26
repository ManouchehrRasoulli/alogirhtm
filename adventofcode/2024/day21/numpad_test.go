package day21

import (
	"testing"
)

func TestNumpad(t *testing.T) {
	numpad := NewNumpad("029A")
	seq, paths := numpad.NumsToDirections()
	for _, path := range paths {
		t.Log(seq, string(path), len(path))
	}
}

func TestPermutations(t *testing.T) {
	runes := []rune{'A', 'B', 'C'}
	result := make(map[string]struct{})
	permute(runes, 0, len(runes)-1, result)

	for k, _ := range result {
		t.Log(k)
	}
}
