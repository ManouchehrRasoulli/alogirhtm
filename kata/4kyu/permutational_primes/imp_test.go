package kata

import "testing"

func Test_generatePrimes(t *testing.T) {
	generatePrimes(2, 20)

	t.Logf("%+v\n", primes)

	generatePrimes(3, 40)

	t.Logf("%+v\n", primes)

	generatePrimes(10, 100)

	t.Logf("%+v\n", primes)

	generatePrimes(10, 1000)

	t.Logf("%+v\n", primes)

	generatePrimes(1, 3000)

	t.Logf("%+v\n", primes)

	generatePrimes(1, 15033)

	t.Logf("%+v\n", primes)
}

func Test_numPermutations(t *testing.T) {
	generatePrimes(1, 4000)

	t.Log(numPermutations(1123, 3000))
}

func Test_PermutationalPrimes(t *testing.T) {
	t.Logf("%+v", PermutationalPrimes(1000, 3))   // [3, 149, 379]
	t.Logf("%+v", PermutationalPrimes(1000, 2))   // [9, 113, 389]
	t.Logf("%+v", PermutationalPrimes(1000, 1))   // [34, 13, 797]
	t.Logf("%+v", PermutationalPrimes(3000, 1))   // [86, 13, 2699]
	t.Logf("%+v", PermutationalPrimes(15033, 7))  // [6, 1123, 4597]
	t.Logf("%+v", PermutationalPrimes(41034, 10)) // [11, 1237, 23473]
	t.Logf("%+v", PermutationalPrimes(5577, 2))   // [47, 113, 4057]
}
