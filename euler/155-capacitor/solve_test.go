package euler

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Log(Solve(1))  // 1
	t.Log(Solve(2))  // 3
	t.Log(Solve(3))  // 7
	t.Log(Solve(4))  // 15
	t.Log(Solve(18)) // 262143
}
