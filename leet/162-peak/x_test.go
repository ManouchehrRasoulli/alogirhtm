package peak

import "testing"

func Test_Peak(t *testing.T) {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	t.Log(findPeakElement(nums))
	t.Log(findPeakElement_(nums))
}
