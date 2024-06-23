package _kyu

import (
	"strings"
)

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	m := make(map[int32]int)
	for _, w := range word {
		m[w] += 1
	}

	s := ""
	for _, w := range word {
		if m[w] > 1 {
			s += ")"
			continue
		}
		s += "("
	}

	return s
}
