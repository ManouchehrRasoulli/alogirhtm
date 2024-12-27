package kata

import (
	"fmt"
	"strconv"
)

func chimesTime(h, m int) int {
	if m == 0 {
		return h
	}

	if m == 15 || m == 30 || m == 45 {
		return 1
	}

	return 0
}

func nextHour(h int) int {
	if h+1 > 12 {
		return 1
	}

	return h + 1
}

func nextTime(h, m int) (int, int) {
	if m < 15 {
		return h, 15
	}

	if m < 30 {
		return h, 30
	}

	if m < 45 {
		return h, 45
	}

	return nextHour(h), 0
}

func CuckooClock(initialTime string, chimes int) string {
	h, _ := strconv.Atoi(initialTime[0:2])
	m, _ := strconv.Atoi(initialTime[3:5])

	for {
		chimes -= chimesTime(h, m)
		if chimes <= 0 {
			break
		}

		h, m = nextTime(h, m)
	}

	return fmt.Sprintf("%02d:%02d", h, m)
}
