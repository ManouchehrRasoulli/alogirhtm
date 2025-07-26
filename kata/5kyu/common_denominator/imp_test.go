package kata

import "testing"

func TestConvertFracts(t *testing.T) {
	t.Log(ConvertFracts([][]int{{1, 2}, {1, 3}, {1, 4}}))
	t.Log(ConvertFracts([][]int{{69, 130}, {87, 1310}, {30, 40}}))
}
