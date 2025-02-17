package kata

import "testing"

func TestFindMult10SF(t *testing.T) {
	t.Log(FindMult10SF(1))  // 60
	t.Log(FindMult10SF(2))  // 70080
	t.Log(FindMult10SF(3))  // 90700800
	t.Log(FindMult10SF(10)) // 556978939118488919905602109440
}
