package kata

import "testing"

func TestKthLastDigitPrime(t *testing.T) {
	t.Log(KthLastDigitPrime(1)) // [92 480150779]
	t.Log(KthLastDigitPrime(2)) // [98 922495169]
	t.Log(KthLastDigitPrime(3)) // [110 603021049]
	t.Log(KthLastDigitPrime(4)) // [122 931883761]
	t.Log(KthLastDigitPrime(5)) // [134 170886697]
	t.Log(KthLastDigitPrime(6)) // [147 167835361]
	t.Log(KthLastDigitPrime(7)) // [200 602829323]
}
