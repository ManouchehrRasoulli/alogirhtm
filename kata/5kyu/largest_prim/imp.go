package kata

import (
	"log"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func BigPrimefacDiv(n int) []int {
	log.Println(n)

	if n < 0 {
		n = -n
	}

	if isPrime(n) {
		return []int{}
	}

	var (
		largestPrime   = 0
		largestDivisor = 0
		num            = n
	)

	for num%2 == 0 {
		largestPrime = 2
		num /= 2
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		for num%i == 0 {
			largestPrime = i
			num /= i
		}
	}

	if num > 2 {
		largestPrime = num
	}

	smp := smallestPrimeFactor(n)
	largestDivisor = n / smp

	return []int{largestPrime, largestDivisor}
}

func smallestPrimeFactor(n int) int {
	if n%2 == 0 {
		return 2
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return i
		}
	}
	return n
}
