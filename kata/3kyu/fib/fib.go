package kata

import (
	"math/big"
)

// Matrix2x2 represents a 2x2 matrix of big.Int values
type Matrix2x2 [2][2]*big.Int

// multiply two 2x2 matrices
func multiply(a, b Matrix2x2) Matrix2x2 {
	result := Matrix2x2{
		{new(big.Int), new(big.Int)},
		{new(big.Int), new(big.Int)},
	}

	result[0][0].Add(new(big.Int).Mul(a[0][0], b[0][0]), new(big.Int).Mul(a[0][1], b[1][0]))
	result[0][1].Add(new(big.Int).Mul(a[0][0], b[0][1]), new(big.Int).Mul(a[0][1], b[1][1]))
	result[1][0].Add(new(big.Int).Mul(a[1][0], b[0][0]), new(big.Int).Mul(a[1][1], b[1][0]))
	result[1][1].Add(new(big.Int).Mul(a[1][0], b[0][1]), new(big.Int).Mul(a[1][1], b[1][1]))

	return result
}

// fast exponentiation of matrix
func matrixPow(matrix Matrix2x2, n int64) Matrix2x2 {
	// identity matrix
	result := Matrix2x2{
		{big.NewInt(1), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1)},
	}

	for n > 0 {
		if n%2 == 1 {
			result = multiply(result, matrix)
		}
		matrix = multiply(matrix, matrix)
		n /= 2
	}

	return result
}

func fib(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}

	isNeg := false
	if n < 0 {
		isNeg = true
		n *= -1
	}

	base := Matrix2x2{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}

	result := matrixPow(base, n-1)
	if !isNeg {
		return result[0][0]
	}

	if n%2 == 0 {
		return new(big.Int).Neg(result[0][0])
	}

	return result[0][0]
}
