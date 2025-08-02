package kata

func ModPow(base, exponent, modulo uint64) uint64 {
	var result uint64 = 1

	for exponent > 0 {
		if exponent%2 == 1 {
			result = result * base % modulo
		}
		exponent /= 2
		base = base * base % modulo
	}

	return result
}
