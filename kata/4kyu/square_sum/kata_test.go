package kata

import "testing"

func TestDecomposeGreedy(t *testing.T) {
	t.Run("11", func(t *testing.T) {
		t.Log(DecomposeGreedy(11))
	})

	t.Run("21", func(t *testing.T) {
		t.Log(DecomposeGreedy(21))
	})

	t.Run("50", func(t *testing.T) {
		t.Log(DecomposeGreedy(50))
	})

	t.Run("7654321", func(t *testing.T) {
		t.Log(DecomposeGreedy(7654321))
	})
}

func TestDecompose(t *testing.T) {
	t.Run("11", func(t *testing.T) {
		t.Log(Decompose(11))
	})

	t.Run("21", func(t *testing.T) {
		t.Log(Decompose(21))
	})

	t.Run("50", func(t *testing.T) {
		t.Log(Decompose(50))
	})

	t.Run("7654321", func(t *testing.T) {
		t.Log(Decompose(7654321))
	})
}
