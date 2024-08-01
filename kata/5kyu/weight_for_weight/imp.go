package kata

import (
	"strings"
)

func sort(param []string, fn func(a, b string) int) {
	for i := 0; i < len(param); i++ {
		for j := i + 1; j < len(param); j++ {
			if fn(param[i], param[j]) > 0 {
				param[i], param[j] = param[j], param[i]
			}
		}
	}
}

func OrderWeight(str string) string {
	sp := strings.Split(str, " ")

	spf := make([]string, 0)
	for _, x := range sp {
		if x == " " {
			continue
		}

		spf = append(spf, x)
	}

	sort(spf, func(a, b string) int {
		sm := func(x string) int {
			xSum := 0
			for _, v := range x {
				xSum += int(v - '0')
			}

			return xSum
		}

		aSum := sm(a)
		bSum := sm(b)

		if aSum == bSum {
			if a <= b {
				return -1
			}
			return 1
		}

		if aSum <= bSum {
			return -1
		}
		return 1
	})

	return strings.Join(spf, " ")
}
