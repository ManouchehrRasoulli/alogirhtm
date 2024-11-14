package kata

var dp = make([]uint64, 1)

func SumOfSquares(n uint64) uint64 {
	if n < uint64(len(dp)-1) {
		return dp[n]
	}

	start := uint64(len(dp) - 1)
	end := n + 1

	for i := start; i <= end; i++ {
		dp = append(dp, n)
		for j := uint64(1); j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}
