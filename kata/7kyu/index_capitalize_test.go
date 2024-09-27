package _kyu

import "testing"

func TestCapitalize(t *testing.T) {
	x := Capitalize("xyz", []int{0, 2})
	t.Log(x)
}
