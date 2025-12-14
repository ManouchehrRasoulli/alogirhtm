package day5

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

func FreshIngredients(input string) (int, int) {
	var (
		lines       = strings.Split(input, "\n")
		ranges      = make([][2]int64, 0)
		ingredients = make([]int64, 0)
		fresh       int
		spoiled     int
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

	for _, id := range ingredients {
		found := false
		for _, rng := range ranges {
			if rng[0] <= id && id <= rng[1] {
				fresh++
				found = true
				break
			}
		}
		if !found {
			spoiled++
		}
	}

	return fresh, spoiled
}
