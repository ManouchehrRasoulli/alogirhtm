package kata

import "testing"

func TestPickPeaks(t *testing.T) {
	v := []int{1, 1, 1, 3, 2, 3, 4, 5, 2, 1}
	t.Log(PickPeaks(v))

	v = []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}
	t.Log(PickPeaks(v))

	v = []int{2, 1, 3, 1, 2, 2, 2, 2, 1}
	t.Log(PickPeaks(v))
}
