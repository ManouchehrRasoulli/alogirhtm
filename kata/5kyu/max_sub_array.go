package _kyu

func MaximumSubarraySum(numbers []int) int {
	var (
		maxSoFar      = 0
		maxEndingHere = 0
	)

	for i := 0; i < len(numbers); i++ {
		maxEndingHere += numbers[i]
		if numbers[i] > maxEndingHere {
			maxEndingHere = numbers[i]
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}

	return maxSoFar
}
