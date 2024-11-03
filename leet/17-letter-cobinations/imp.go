package lettercobinations

var keys map[byte][]string = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var resault []string

	var backTrace func(digits string, loc int, current string)

	backTrace = func(digits string, loc int, current string) {
		if len(current) != 0 && loc == len(digits) {
			resault = append(resault, current)
			return
		}

		for _, i := range keys[digits[loc]] {
			backTrace(digits, loc+1, current+i)
		}
	}

	if len(digits) != 0 {
		backTrace(digits, 0, "")
	}

	return resault
}
