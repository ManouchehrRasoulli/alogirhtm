package groupanagrams

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, s := range strs {
		sb := []rune(s)
		slices.Sort[[]rune](sb)
		ss := string(sb)
		if _, ok := m[ss]; !ok {
			m[ss] = []string{s}
			continue
		}

		m[ss] = append(m[ss], s)
	}

	re := make([][]string, 0)
	for _, v := range m {
		re = append(re, v)
	}

	return re
}
