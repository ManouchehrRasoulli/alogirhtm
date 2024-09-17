package kata

import "testing"

func TestSceene(t *testing.T) {
	var boar = [4][4]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4},
	}

	t.Log(scenes(0, 0, right, boar))
	t.Log(scenes(0, 3, left, boar))
	t.Log(scenes(0, 3, down, boar))
	t.Log(scenes(3, 3, up, boar))
	t.Log(scenes(3, 2, up, boar))
	t.Log(scenes(2, 0, right, boar))
}

func TestFindClueLocation(t *testing.T) {
	t.Log(getClues(1, 2))
	SolvePuzzle([]int{})
}
