package kata

import "testing"

func TestFold(t *testing.T) {
	t.Log(fold([]int{1, 2, 3}))          // [4, 2]
	t.Log(fold([]int{1, 2, 3, 4}))       // [5, 5]
	t.Log(fold([]int{1, 2, 3, 4, 5}))    // [6, 6, 3]
	t.Log(fold([]int{1, 2, 3, 4, 5, 6})) // [7, 7, 7]
}

func TestFoldArray(t *testing.T) {
	t.Log(FoldArray([]int{1, 2, 3, 4, 5}, 1)) // [6, 6, 3]
	t.Log(FoldArray([]int{1, 2, 3, 4, 5}, 2)) // [9, 6]
	t.Log(FoldArray([]int{1, 2, 3, 4, 5}, 3)) // [15]
	t.Log(FoldArray([]int{1, 2, 3, 4, 5}, 4)) // [15]
}
