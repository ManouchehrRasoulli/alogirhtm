package kata

import "testing"

func Test_permEq(t *testing.T) {
	t.Log(permEq(127, 712)) // true
	t.Log(permEq(127, 713)) // false
	t.Log(permEq(133, 931)) // false
}

func Test_SearchPermMult(t *testing.T) {
	t.Log(SearchPermMult(10000, 7)) // == 1 (1359, 9513)
	t.Log(SearchPermMult(10000, 4)) // == 2 (1782, 7128) and (2178, 8712)
}
