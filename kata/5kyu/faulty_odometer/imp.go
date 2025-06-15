package kata

import (
	"strconv"
	"strings"
)

func FaultyOdometerNaive(n int) int {
	var (
		count        = 0
		containsFour = func(i int) bool {
			return strings.Contains(strconv.Itoa(i), "4")
		}
	)

	for i := 1; i <= n; i++ {
		if containsFour(i) {
			continue
		}

		count++
	}

	return count
}
