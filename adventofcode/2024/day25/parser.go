package day25

const (
	value     = '#'
	empty     = '.'
	allValue  = "#####"
	maxHeight = 6
)

func ReadCombinations(lines []string) ([]*Key, []*Lock) {
	// read in all combinations and data
	var (
		keys  = make([]*Key, 0)
		locks = make([]*Lock, 0)
	)

	for i := 0; i < len(lines); i += 8 {
		combination := readIn(lines[i : i+7])

		if lines[i] == allValue { // it's a lock
			locks = append(locks, NewLock(combination))
		} else if lines[i+6] == allValue { // it's a key
			keys = append(keys, NewKey(combination))
		}
	}

	return keys, locks
}

func readIn(s []string) []int {
	var (
		data = make([][]int, 0)
	)
	for i := 0; i < 6; i++ {
		line := strToInt(s[i])
		data = append(data, line)
	}

	return zip(data)
}

func strToInt(s string) []int {
	var (
		values = make([]int, 0)
	)

	for _, v := range s {
		switch v {
		case empty:
			values = append(values, 0)
		case value:
			values = append(values, 1)
		}
	}

	return values
}

func zip(data [][]int) []int {
	if len(data) == 0 {
		return []int{}
	}

	var (
		zp = make([]int, len(data[0]))
	)

	for _, v := range data {
		for j, jv := range v {
			zp[j] += jv
		}
	}

	return zp
}
