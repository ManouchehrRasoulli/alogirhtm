package _kyu

func TrailingZeros(n int) int {
	zeros := 0

	for n > 0 {
		zeros += n / 5
		n /= 5
	}

	return zeros
}
