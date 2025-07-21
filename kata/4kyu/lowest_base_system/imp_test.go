package kata

import "testing"

func TestGetMinBase(t *testing.T) {
	t.Log(GetMinBase(7))               // 2
	t.Log(GetMinBase(57))              // 7
	t.Log(GetMinBase(125002500050001)) // 50000
	t.Log(GetMinBase(1000000000000))   // 999999999999
	t.Log(GetMinBase(1001001))         // 1000
}
