package kata

import (
	"math"
	"strconv"
)

var (
	primes = []int{2}
)

func _min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func _max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func isInNumbers(num int, nums []int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}

	return false
}

func numPermutations(num int, maxPrime int) []int {
	var res []int
	numStr := strconv.Itoa(num)

	generatePermutations([]rune(numStr), 0, &res)

	// store generated data
	var resPrimes []int
	for _, v := range res {
		if v == num {
			continue // ignore the number
		}

		if v > maxPrime {
			continue
		}

		ok := isInNumbers(v, primes)
		if !ok {
			// number is not prime
			continue
		}

		ok = isInNumbers(v, resPrimes)
		if ok {
			continue
		}

		resPrimes = append(resPrimes, v)
	}

	return resPrimes
}

func generatePermutations(num []rune, start int, res *[]int) {
	if start == len(num)-1 {
		if num[0] == '0' {
			return
		}

		numStr := string(num)
		num, _ := strconv.Atoi(numStr)
		*res = append(*res, num)
	}

	for i := start; i < len(num); i++ {
		num[start], num[i] = num[i], num[start] // swap out
		generatePermutations(num, start+1, res)
		num[start], num[i] = num[i], num[start] // swap back
	}
}

func generatePrimes(x, y int) {
	if x > y {
		x, y = y, x
	}

	// no data to generate
	if y < 2 {
		return
	}

	if y < primes[len(primes)-1] {
		// already generated
		return
	}

	if x <= primes[len(primes)-1] {
		x = primes[len(primes)-1] + 1
	}

	sieve := make([]bool, y+1)
	for i := 2; i*i <= y; i++ {
		if !sieve[i] {
			for j := i * i; j <= y; j += i {
				sieve[j] = true
			}
		}
	}

	rangePrimes := make([]int, 0)
	for i := x; i <= y; i++ {
		if !sieve[i] {
			rangePrimes = append(rangePrimes, i)
		}
	}

	primes = append(primes, rangePrimes...)
}

func PermutationalPrimes(nMax, kPerms int) [3]int {
	maxPrime := primes[len(primes)-1]
	if nMax > maxPrime {
		generatePrimes(2, nMax)
	}

	var (
		count         = 0
		minValidPrime = math.MaxInt64
		maxValidPrime = math.MinInt64
	)

	foundPermutations := make(map[int]struct{})

	for i := 0; i < len(primes); i++ {
		if primes[i] > nMax {
			break
		}

		primeKPerms := numPermutations(primes[i], nMax)

		if _, alreadyComputed := foundPermutations[primes[i]]; alreadyComputed {
			continue
		}

		if len(primeKPerms) == kPerms {
			for _, primeKPerm := range primeKPerms {
				foundPermutations[primeKPerm] = struct{}{}
			}

			foundPermutations[primes[i]] = struct{}{}

			count++
			minValidPrime = _min(primes[i], minValidPrime)
			maxValidPrime = _max(primes[i], maxValidPrime)
		}
	}

	if count == 0 {
		minValidPrime = 0
		maxValidPrime = 0
	}

	return [3]int{count, minValidPrime, maxValidPrime}
}
