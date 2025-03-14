package day7

func GenerateCombinationOfOperationInLen(operators []string, x int, current string) []string {
	if x == 0 {
		return []string{current}
	}

	var result []string

	for _, op := range operators {
		result = append(result, GenerateCombinationOfOperationInLen(operators, x-1, current+op)...)
	}

	return result
}
