package kata

func min_(a, b, c uint) uint {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

var hamming = []uint{1}

func Hammer(n int) uint {
	if n < len(hamming) {
		return hamming[n-1]
	}

	v := len(hamming) - 1
	i2, i3, i5 := v, v, v

	for i := len(hamming); i <= n; i++ {
		nextHamming := min_(hamming[i2]*2, hamming[i3]*3, hamming[i5]*5)
		hamming = append(hamming, nextHamming)

		if nextHamming == hamming[i2]*2 {
			i2 += 1
		}

		if nextHamming == hamming[i3]*3 {
			i3 += 1
		}

		if nextHamming == hamming[i5]*5 {
			i5 += 1
		}
	}

	return hamming[n-1]
}
