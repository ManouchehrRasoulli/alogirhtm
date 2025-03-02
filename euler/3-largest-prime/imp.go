package __largest_prime

import "math"

func isPrime(num uint64) bool {
	if num == 2 || num == 3 {
		return true
	}

	if num%2 == 0 || num%3 == 0 {
		return false
	}

	numSqrt := uint64(math.Sqrt(float64(num)))
	for i := uint64(5); i <= numSqrt; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func nextPrime(current uint64) uint64 {
	x := current + 1
	for {
		if isPrime(x) {
			return x
		}
		x++
	}
}

func LargestPrimeFactor(n uint64) uint64 {
	largest := uint64(2)

	for n != 1 {
		if n%largest == 0 {
			n /= largest
		} else {
			largest = nextPrime(largest)
		}
	}

	return largest
}
