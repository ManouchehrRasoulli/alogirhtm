package kata

import (
	"math"
	"math/big"
)

func generatePermutations(num []rune, start int, smallest, biggest *big.Int, data map[string]struct{}) {
	if start == len(num)-1 {
		num, _ := big.NewInt(0).SetString(string(num), 10)

		if _, ok := data[num.Text(10)]; ok {
			return
		}

		if num.Cmp(smallest) == -1 {
			*smallest = *num
		}

		if num.Cmp(biggest) == 1 {
			*biggest = *num
		}

		data[num.Text(10)] = struct{}{}

		return
	}

	for i := start; i < len(num); i++ {
		num[start], num[i] = num[i], num[start] // swap out
		generatePermutations(num, start+1, smallest, biggest, data)
		num[start], num[i] = num[i], num[start] // swap back
	}
}

func ProcArr(num []rune) [3]*big.Int {
	data := make(map[string]struct{})
	smallest := big.NewInt(math.MaxInt64)
	biggest := big.NewInt(-math.MaxInt64)

	generatePermutations(num, 0, smallest, biggest, data)

	return [3]*big.Int{big.NewInt(int64(len(data))),
		smallest, biggest}
}
