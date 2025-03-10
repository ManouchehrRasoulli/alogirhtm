package kata

import "math"

func FindNb(m int) int {
	return volume(m, 0, 0)
}

func volume(m int, n int, cv int) int {
	if cv == m {
		return n - 1
	}

	if cv > m {
		return -1
	}

	cv += int(math.Pow(float64(n), 3))
	return volume(m, n+1, cv)
}
