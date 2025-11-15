package kata

import "math"

func F(n int) int {
	var (
		result = 1
		dv     []int
	)
	dv = derivation(factors(n))
	for _, v := range dv {
		result *= v
	}

	return result
}

func factors(n int) map[int]int {
	fac := make(map[int]int)

	for i := 2; i <= n; i++ {
		for n%i == 0 && n > 0 {
			fac[i]++
			n /= i
		}
	}

	return fac
}

func derivation(fac map[int]int) []int {
	var (
		result = make([]int, 0)
		cmp    = func(k, v int) int {
			x := math.Pow(float64(k), float64(v-1))
			return v * int(x)
		}
	)

	for k, v := range fac {
		result = append(result, cmp(k, v))
	}

	return result
}
