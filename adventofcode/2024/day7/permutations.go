package day7

var (
	operators = []string{"+", "*"}
)

func GenerateCombinationOfOperationInLen(x int, current string) []string {
	if x == 0 {
		return []string{current}
	}

	var result []string

	for _, op := range operators {
		result = append(result, GenerateCombinationOfOperationInLen(x-1, current+op)...)
	}

	return result
}
