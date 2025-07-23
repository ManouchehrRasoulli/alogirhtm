package kata

import "fmt"

func lcsLen(x, y string) [][]int {
	m, n := len(x), len(y)
	matrix := make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				if matrix[i-1][j] > matrix[i][j-1] {
					matrix[i][j] = matrix[i-1][j]
				} else {
					matrix[i][j] = matrix[i][j-1]
				}
			}
		}
	}
	return matrix
}

func backtrack(matrix [][]int, x, y string) string {
	var (
		i, j  = len(x), len(y)
		chars = make([]byte, 0)
	)

	for i > 0 && j > 0 {
		if x[i-1] == y[j-1] {
			chars = append([]byte{x[i-1]}, chars...)
			i--
			j--
		} else if matrix[i-1][j] > matrix[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(chars)
}

func LCS(x, y string) string {
	matrix := lcsLen(x, y)

	xNone := fmt.Sprintf("#%s", x)
	yNone := fmt.Sprintf("#%s", y)
	fmt.Println(yNone)
	for i := 0; i < len(xNone); i++ {
		fmt.Println(string(xNone[i]), matrix[i])
	}
	return backtrack(matrix, x, y)
}
