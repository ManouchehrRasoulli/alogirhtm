package regular_expression

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// Create two slices to represent the DP table
	dp := make([]bool, n+1)
	prev := make([]bool, n+1)

	// Empty pattern matches empty string
	dp[0] = true

	// Handle patterns like a* or a*b* at the beginning
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[j] = dp[j-2]
		}
	}

	// Build the DP table
	for i := 1; i <= m; i++ {
		// Swap slices to reuse the space
		dp, prev = prev, dp
		// Reset the current row
		dp[0] = false

		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[j] = prev[j-1]
			} else if p[j-1] == '*' {
				// Case 1: '*' acts as zero characters
				dp[j] = dp[j-2]
				// Case 2: '*' acts as one or more characters
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[j] = dp[j] || prev[j]
				}
			} else {
				dp[j] = false
			}
		}
	}

	return dp[n]
}
