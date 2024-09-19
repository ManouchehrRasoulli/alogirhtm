package kata

func binomial(n int, k int) uint64 {
	if k == 0 {
		return 1
	}

	if k == 1 {
		return uint64(n)
	}

	return uint64(n) * binomial(n-1, k-1) / uint64(k)
}

func TotalIncDec(n int) uint64 {
	inc := binomial(n+9, n)
	dec := binomial(n+10, n)
	dup := uint64(10*n + 1)

	return inc + dec - dup
}
