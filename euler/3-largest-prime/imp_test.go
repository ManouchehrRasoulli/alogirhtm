package __largest_prime

import "testing"

func TestNextPrime(t *testing.T) {
	v := nextPrime(3)
	t.Log(v) // 5

	v = nextPrime(1000)
	t.Log(v) // 1009
}

func TestLargestPrime(t *testing.T) {
	x := LargestPrimeFactor(13195)
	t.Log(x) // 29

	x = LargestPrimeFactor(600851475143)
	t.Log(x)
}
