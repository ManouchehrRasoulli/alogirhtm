package findminrotated

import "testing"

func Test_findMin(t *testing.T) {
	t.Log(findMin([]int{3, 4, 5, 1, 2}))
	t.Log(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	t.Log(findMin([]int{11, 13, 15, 17}))
	t.Log(findMin([]int{1}))
}
