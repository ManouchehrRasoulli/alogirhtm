package peak

import "testing"

func Test_Peak(t *testing.T) {
	t.Log(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))

	t.Log(findPeakElement([]int{-2147483648, -2147483647}))
}
