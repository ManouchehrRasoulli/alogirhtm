package _kyu

import "testing"

func Test_FindHeight(t *testing.T) {
	t.Log(FindHeight(0))     // 0
	t.Log(FindHeight(1))     // 1
	t.Log(FindHeight(2))     // 1
	t.Log(FindHeight(3))     // 1
	t.Log(FindHeight(4))     // 2
	t.Log(FindHeight(5))     // 2
	t.Log(FindHeight(10))    // 3
	t.Log(FindHeight(43419)) // 62
}
