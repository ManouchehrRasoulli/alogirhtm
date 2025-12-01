package kata

import "testing"

func TestProperFractions(t *testing.T) {
	t.Log(ProperFractions(1) == 0)
	t.Log(ProperFractions(2) == 1)
	t.Log(ProperFractions(5) == 4)
	t.Log(ProperFractions(15) == 8)
	t.Log(ProperFractions(25) == 20)
	t.Log(ProperFractions(9999999) == 6637344)
	t.Log(ProperFractions(500000003) == 500000002)
}
