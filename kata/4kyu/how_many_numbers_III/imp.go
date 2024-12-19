package kata

var (
	count int
	min   int
	max   int
)

func FindAll(sumDigits, numDigits int) []int {
	count = 0
	min = 0
	max = 0

	RecursiveSearch(0, 1, sumDigits, numDigits)

	var result []int
	if count > 0 {
		result = append(result, count, min, max)
	}

	return result
}

func RecursiveSearch(currNum, prevDigit, sumLeft, digitsLeft int) {
	if sumLeft == 0 && digitsLeft == 0 {
		if count == 0 {
			min = int(currNum)
		}
		if min > int(currNum) {
			min = int(currNum)
		}
		if max < int(currNum) {
			max = int(currNum)
		}
		count++
	} else if digitsLeft != 0 {
		for i := prevDigit; i < 10; i++ {
			RecursiveSearch(10*currNum+i, i, sumLeft-i, digitsLeft-1)
		}
	}
}
