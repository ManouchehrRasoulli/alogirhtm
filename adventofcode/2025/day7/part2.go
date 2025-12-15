package day7

import (
	"strings"
)

func ManifoldDimensions(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	h := len(lines)
	w := len(lines[0])

	field := make([][]string, h)
	for i := range field {
		field[i] = strings.Split(lines[i], "")
	}

	timelines := make([][]int64, h)
	for i := range timelines {
		timelines[i] = make([]int64, w)
	}

	for j := 0; j < w; j++ {
		if field[0][j] == Start {
			timelines[0][j] = 1
		}
	}

	for i := 1; i < h; i++ {
		prev := timelines[i-1]
		curr := timelines[i]

		for j := 0; j < w; j++ {
			if prev[j] == 0 {
				continue
			}

			switch field[i][j] {
			case Split:
				if j-1 >= 0 {
					curr[j-1] += prev[j]
				}
				if j+1 < w {
					curr[j+1] += prev[j]
				}
			case Empty:
				curr[j] += prev[j]
			}
		}
	}

	var total int64
	for _, v := range timelines[h-1] {
		total += v
	}

	return int(total)
}
