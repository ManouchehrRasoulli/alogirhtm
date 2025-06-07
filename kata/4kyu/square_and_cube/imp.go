package kata

import (
	"fmt"
	"math/big"
	"strconv"
)

func reverse(num int) int {
	var (
		reverseNumStr = ""
		numStr        = fmt.Sprintf("%d", num)
	)

	for i := len(numStr) - 1; i >= 0; i-- {
		reverseNumStr += string(numStr[i])
	}

	reverseNum, _ := strconv.Atoi(reverseNumStr)
	return reverseNum
}

var (
	memo = make([]int, 0)
)

func SqCubRevPrime(n int) int {
	if len(memo) >= n {
		return memo[n-1]
	}

	count := len(memo)
	i := 1
	if count > 0 {
		i = memo[len(memo)-1] + 1
	}

	for ; count < n; i++ {
		sq := i * i
		cub := i * i * i
		if big.NewInt(int64(reverse(sq))).ProbablyPrime(20) &&
			big.NewInt(int64(reverse(cub))).ProbablyPrime(20) {

			memo = append(memo, i)
			count++
		}
	}
	return memo[n-1]
}
