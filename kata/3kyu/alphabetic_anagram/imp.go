package kata

import (
	"math/big"
)

func ListPosition(word string) *big.Int {
	var (
		rank = big.NewInt(1)
		n    = len(word)
		freq = make([]int, 26)
		fact = make([]*big.Int, n+1)
	)

	for _, ch := range word {
		freq[ch-'A']++
	}

	fact[0] = big.NewInt(1)
	for i := 1; i <= n; i++ {
		fact[i] = new(big.Int).Mul(fact[i-1], big.NewInt(int64(i)))
	}

	for i := 0; i < n; i++ {
		ch := word[i]

		for c := byte('A'); c < ch; c++ {
			if freq[c-'A'] == 0 {
				continue
			}

			freq[c-'A']--
			permutations := new(big.Int).Set(fact[n-i-1])
			denom := big.NewInt(1)
			for j := 0; j < 26; j++ {
				if freq[j] > 1 {
					denom.Mul(denom, fact[freq[j]])
				}
			}
			permutations.Div(permutations, denom)
			rank.Add(rank, permutations)

			freq[c-'A']++
		}

		freq[ch-'A']--
	}

	return rank
}
