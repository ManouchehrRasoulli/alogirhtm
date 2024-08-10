package kata

import "testing"

func TestDownSum(t *testing.T) {
	pyramid := [][]int{{3}, {7, 4}, {2, 4, 6}, {8, 5, 9, 3}}
	t.Log(LongestSlideDown(pyramid))
	t.Log(pyramid)
}
