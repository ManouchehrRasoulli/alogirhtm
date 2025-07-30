package kata

func FindMissingNumber(arr []int) int {
	n := len(arr)
	if n < 2 {
		return -1
	}

	minVal, maxVal := arr[0], arr[0]
	sum := 0

	for _, num := range arr {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
		sum += num
	}

	expectedSum := (n + 1) * (minVal + maxVal) / 2

	return expectedSum - sum
}
