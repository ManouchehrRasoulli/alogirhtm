package _kyu

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func score(s string) int {
	sum := 0
	for _, c := range s {
		x := fmt.Sprintf("%d", c)
		xi, _ := strconv.Atoi(x)
		sum += xi - 96
	}

	return sum
}

func HighWithSort(s string) string {
	words := strings.Split(s, " ")
	slices.SortFunc(words, func(a, b string) int {
		aScore, bScore := score(a), score(b)
		if aScore > bScore {
			return -1
		}
		if aScore == bScore {
			return 0
		}
		return 1
	})
	return words[0]
}

func High(s string) string {
	words := strings.Split(s, " ")
	var (
		maxScore int
		maxWord  string
	)
	for _, w := range words {
		wScore := score(w)
		if wScore > maxScore {
			maxScore = wScore
			maxWord = w
		}
	}
	return maxWord
}
