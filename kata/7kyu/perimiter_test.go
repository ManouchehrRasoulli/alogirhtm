package _kyu

import "testing"

func TestEquableTriangle(t *testing.T) {
	t.Log(EquableTriangle(0, 0, 0))
	t.Log(EquableTriangle(2, 3, 4))
	t.Log(EquableTriangle(12, 47, 21))
	t.Log(EquableTriangle(6, 55, 93))
	t.Log(EquableTriangle(41, 97, 57))
}
