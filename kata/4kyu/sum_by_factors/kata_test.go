package kata

import "testing"

func Test_isPrime(t *testing.T) {
	t.Log(isPrime(3))
	t.Log(isPrime(7))
	t.Log(isPrime(9))
	t.Log(isPrime(11))
	t.Log(isPrime(13))
	t.Log(isPrime(15))
	t.Log(isPrime(17))
}

func Test_SumOfDivided1(t *testing.T) {
	t.Log(SumOfDivided([]int{12, 15}))
	t.Log(SumOfDivided([]int{15, 21, 24, 30, 45}))
	t.Log(SumOfDivided([]int{-29804, -4209, -28265, -72769, -31744}))
}
