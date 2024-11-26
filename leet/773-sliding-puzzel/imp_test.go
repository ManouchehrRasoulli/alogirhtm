package kata

import "testing"

func Test_slidingPuzzle(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 0, 5},
	}

	t.Log(slidingPuzzle(board))
}
