package kata

import (
	"math"
	"math/big"
	"slices"
)

func greedySearch(remaining int64, candidate int64) []int64 {
	if remaining == 0 {
		return []int64{}
	}
	if candidate == 0 {
		return nil
	}
	if candidate*candidate > remaining {
		return greedySearch(remaining, candidate-1)
	}
	result := greedySearch(remaining-(candidate*candidate), candidate-1)
	if result != nil {
		return append([]int64{candidate}, result...)
	}
	return greedySearch(remaining, candidate-1)
}

func DecomposeGreedy(n int64) []int64 {
	result := greedySearch(n*n, n-1)
	slices.Reverse(result)
	return result
}

func calculateMaxSum(candidate int64) *big.Int {
	c := big.NewInt(candidate)
	c1 := new(big.Int).Add(c, big.NewInt(1))
	c2 := new(big.Int).Mul(big.NewInt(2), c)
	c2.Add(c2, big.NewInt(1))

	result := new(big.Int).Mul(c, c1)
	result.Mul(result, c2)
	result.Div(result, big.NewInt(6))

	return result
}

func backtrackSearch(remaining int64, candidate int64) []int64 {
	if remaining == 0 {
		return []int64{}
	}

	if candidate <= 0 {
		return nil
	}

	sqrt := int64(math.Sqrt(float64(remaining)))
	candidate = min(candidate, sqrt)

	if candidate <= 0 {
		return nil
	}

	maxSum := calculateMaxSum(candidate)
	remainingBig := big.NewInt(remaining)

	if maxSum.Cmp(remainingBig) < 0 {
		return nil
	}

	chosen := backtrackSearch(
		remaining-candidate*candidate,
		candidate-1,
	)

	if chosen != nil {
		return append([]int64{candidate}, chosen...)
	}

	return backtrackSearch(remaining, candidate-1)
}

func loop(s int64, i int64) []int64 {
	if s < 0 {
		return nil
	}
	if s == 0 {
		return []int64{}
	}
	var j int64 = i - 1
	for ; j > 0; j-- {
		var sub = loop(s-j*j, j)
		if sub != nil {
			return append(sub, []int64{j}...)
		}
	}
	return nil
}

func Decompose(n int64) []int64 {
	return loop(n*n, n)
}
