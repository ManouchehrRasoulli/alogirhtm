package rotate_image

import "testing"

func TestRotateImage(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	rotate(m)

	t.Log(m)
}
