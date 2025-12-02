package kata

func ProperFractions(n int) int {
	if n <= 1 {
		return 0
	}

	result := n
	d := n

	for p := 2; p*p <= d; p++ {
		if d%p == 0 {
			result -= result / p
			for d%p == 0 {
				d /= p
			}
		}
	}

	if d > 1 {
		result -= result / d
	}

	return result
}
