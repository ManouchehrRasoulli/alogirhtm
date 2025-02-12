package kata

import "testing"

func TestFindOdd(t *testing.T) {
	t.Log(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5})) // 5
	t.Log(FindOdd([]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}))                      // -1
	t.Log(FindOdd([]int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}))                        // 10
}
