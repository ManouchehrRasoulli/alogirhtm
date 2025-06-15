package kata

import (
	"strconv"
	"strings"
)

func FaultyOdometerNaive(n int) int {
	var (
		count        = 0
		containsFour = func(i int) bool {
			return strings.Contains(strconv.Itoa(i), "4")
		}
	)

	for i := 1; i <= n; i++ {
		if containsFour(i) {
			continue
		}

		count++
	}

	return count
}

func FaultyOdometerNaiveSecond(n int) int {
	var (
		count        = 0
		containsFour = func(i int) int {
			iStr := strconv.Itoa(i)
			v := strings.Index(iStr, "4")
			if v < 0 {
				return v
			}

			return len(iStr) - v - 1
		}
	)

	for i := 1; i <= n; {
		if index := containsFour(i); index >= 0 {
			next := i + index*10
			if next == i {
				next += 1
			}
			i = next
			continue
		}

		i++
		count++
	}

	return count
}

func FaultyOdometer(n int) int {
	digitMap := map[rune]rune{
		'0': '0', '1': '1', '2': '2', '3': '3',
		'5': '4', '6': '5', '7': '6', '8': '7', '9': '8',
	}

	nStr := strconv.Itoa(n)
	base9Str := ""

	for _, ch := range nStr {
		mapped, ok := digitMap[ch]
		if !ok {
			// skip invalid digits like '4'
			continue
		}
		base9Str += string(mapped)
	}

	// Parse base-9 number as base-10 integer
	result, err := strconv.ParseInt(base9Str, 9, 64)
	if err != nil {
		panic(err)
	}

	return int(result)
}
