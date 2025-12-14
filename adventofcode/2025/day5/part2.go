package day5

import (
	"fmt"
	"log"
	"math/big"
	"slices"
	"strings"
)

func FreshIngredientRanges(input string) string {
	var (
		lines       = strings.Split(input, "\n")
		ranges      = make([][2]int64, 0)
		ingredients = make([]int64, 0)
		mergeFunc   = func(ranges [][2]int64) [][2]int64 {
			if len(ranges) == 1 {
				return ranges
			}
			var merged [][2]int64

			current := ranges[0]
			for i := 1; i < len(ranges); i++ {
				if ranges[i][0] <= current[1] {
					current[1] = max(current[1], ranges[i][1])
					continue
				}

				merged = append(merged, current)
				current = ranges[i]
			}

			merged = append(merged, current)

			return merged
		}
		fresh = new(big.Int)
	)

	for _, line := range lines {
		if strings.Contains(line, "-") {
			var start, end int64
			_, err := fmt.Sscanf(line, "%d-%d", &start, &end)
			if err != nil {
				log.Fatalf("%s, %s", err, line)
			}
			ranges = append(ranges, [2]int64{start, end})
			continue
		}

		if len(line) != 0 {
			var ingredient int64
			_, err := fmt.Sscanf(line, "%d", &ingredient)
			if err != nil {
				log.Fatalf("%s, %s", err, line)
			}
			ingredients = append(ingredients, ingredient)
		}
	}

	slices.SortFunc(ranges, func(a, b [2]int64) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		if a[1] < b[1] {
			return -1
		}
		if a[1] > b[1] {
			return 1
		}
		return 0
	})

	ranges = mergeFunc(ranges)
	for _, rg := range ranges {
		fresh = fresh.Add(fresh, big.NewInt(rg[1]-rg[0]+1))
	}

	return fresh.String()
}
