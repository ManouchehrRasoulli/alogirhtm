package distinct

func solve(mx [][]int, cs, ct int, s, t string) int {
	if ct >= len(t) {
		return 1
	}

	if cs >= len(s) {
		return 0
	}

	if mx[cs][ct] != -1 {
		return mx[cs][ct]
	}

	res := solve(mx, cs+1, ct, s, t)
	if s[cs] == t[ct] {
		res += solve(mx, cs+1, ct+1, s, t)
	}

	mx[cs][ct] = res
	return res
}

func numDistinct(s string, t string) int {
	matrix := make([][]int, len(s))
	for i := range matrix {
		matrix[i] = make([]int, len(t))
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	return solve(matrix, 0, 0, s, t)
}
