package kata

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	c := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			c += 2
		}
	}

	if c == 2 {
		return true
	}

	return false
}

func SumOfDivided(lst []int) string {
	mx := math.MinInt
	for i := 0; i < len(lst); i++ {
		mx = int(math.Max(float64(mx), math.Abs(float64(lst[i]))))
	}

	primes := make([]int, 0)
	for i := 2; i <= mx; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	str := ""
	for i := 0; i < len(primes); i++ {
		primeDevides := false
		sum := 0
		for j := 0; j < len(lst); j++ {
			if lst[j]%primes[i] == 0 {
				sum += lst[j]
				primeDevides = true
			}
		}
		if primeDevides {
			str = fmt.Sprintf("%s(%d %d)", str, primes[i], sum)
		}
	}

	return str
}
