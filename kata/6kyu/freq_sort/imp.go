package kata

import (
	"slices"
)

func Solve(arr []int) []int {
	frq := make(map[int]int)

	for _, v := range arr {
		frq[v] += 1
	}

	slices.SortFunc[[]int](arr, func(a, b int) int {
		if frq[a] > frq[b] {
			return -1
		}

		if frq[b] > frq[a] {
			return 1
		}

		if a > b {
			return -1
		}

		return 1
	})

	return arr
}
