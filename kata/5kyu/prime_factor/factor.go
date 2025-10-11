package kata

func getUniquePrimeFactors(n int) []int {
	if n <= 1 {
		return []int{}
	}

	factors := make(map[int]bool)

	if n%2 == 0 {
		factors[2] = true
		for n%2 == 0 {
			n /= 2
		}
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			factors[i] = true
			for n%i == 0 {
				n /= i
			}
		}
	}

	if n > 1 {
		factors[n] = true
	}

	result := make([]int, 0, len(factors))
	for k := range factors {
		result = append(result, k)
	}

	return result
}

func reverse(n int) int {
	result := 0
	for n > 0 {
		result = result*10 + n%10
		n /= 10
	}
	return result
}

func haveSameFactors(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return false
	}

	aMap := make(map[int]bool, len(a))
	for _, v := range a {
		aMap[v] = true
	}

	for _, v := range b {
		if !aMap[v] {
			return false
		}
	}

	return true
}

func SameFactRev(n int) []int {
	result := make([]int, 0)

	for i := 1089; i < n; i++ {
		rev := reverse(i)
		if rev == i {
			continue
		}

		factorsI := getUniquePrimeFactors(i)
		factorsRev := getUniquePrimeFactors(rev)

		if haveSameFactors(factorsI, factorsRev) {
			result = append(result, i)
		}
	}

	return result
}
