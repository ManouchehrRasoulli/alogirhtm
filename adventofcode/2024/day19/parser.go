package day19

import (
	"slices"
	"strings"
)

func ParseInput(data string) ([]string, []string) {
	var (
		patterns []string
		towels   []string
	)

	lines := strings.Split(data, "\n")
	for i, line := range lines {
		switch i {
		case 0:
			items := strings.Split(line, ",")
			for _, item := range items {
				item = strings.TrimSpace(item)
				patterns = append(patterns, item)
			}
		default:
			if len(line) == 0 {
				continue
			}
			towels = append(towels, line)
		}
	}

	slices.SortFunc(patterns, func(a, b string) int {
		if len(a) < len(b) {
			return 1
		}
		return -1
	})
	return patterns, towels
}
