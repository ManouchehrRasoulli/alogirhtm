package kata

import "testing"

func TestPrimeFactors(t *testing.T) {
	v := 2178
	t.Logf("v is %d, reverse is %d", v, reverse(v))
	factors := getUniquePrimeFactors(v)
	t.Logf("factors: %v", factors)

	factors = getUniquePrimeFactors(reverse(v))
	t.Logf("factors: %v", factors)
}

func TestSameFactRev(t *testing.T) {
	t.Logf("%+v\n", SameFactRev(1100))
	t.Logf("%+v\n", SameFactRev(2500))
	t.Logf("%+v\n", SameFactRev(5000))
	t.Logf("%+v\n", SameFactRev(14013))
	t.Logf("%+v\n", SameFactRev(49240))
	t.Logf("%+v\n", SameFactRev(42980))
	t.Logf("%+v\n", SameFactRev(100000))
}
