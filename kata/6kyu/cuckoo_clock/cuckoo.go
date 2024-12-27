package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func toString(h, m int) string {
	return fmt.Sprintf("%02d:%02d", h, m)
}

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
	hs, ms := strings.Split(initialTime, ":")[0], strings.Split(initialTime, ":")[1]
	h, _ := strconv.Atoi(hs)
	m, _ := strconv.Atoi(ms)

	chimesNow := 0
	for {
		chimesNow += chimesTime(h, m)
		if chimesNow >= chimes {
			break
		}

		h, m = nextTime(h, m)
	}

	return toString(h, m)
}
