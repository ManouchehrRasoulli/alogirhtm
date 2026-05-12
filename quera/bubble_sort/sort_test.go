package bubble_sort

import "testing"

func TestSort(t *testing.T) {
	list := []int{3, 2, 4, 1, 5}
	t.Log(list)
	sort(list)
	t.Log(list)
}
