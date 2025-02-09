package kata

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func joinInts(arr []int, sep string) string {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = strconv.Itoa(num)
	}
	return strings.Join(strArr, sep)
}

// repair arithmetic sequence
func repairSequence(arr []int) []int {
	if len(arr) < 3 {
		return arr
	}

	fmt.Printf("input :: {%s}\n", joinInts(arr, ","))

	commonDifference := make(map[int]int)

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < 0 {
			diff = -diff
		}

		commonDifference[diff]++
	}

	fmt.Println(commonDifference)

	if len(commonDifference) > 3 {
		return arr
	}

	var (
		maxDiff = math.MinInt
		maxVal  = 0
	)

	for k, diff := range commonDifference {
		if diff > maxDiff {
			maxDiff = diff
			maxVal = k
		}
	}

	fmt.Println(maxVal, maxDiff)

	// repair
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff != maxVal && i == len(arr)-1 {
			arr[i] = arr[i-1] + maxVal
			break
		}

		if diff != maxVal && i-1 == 0 {
			if arr[i-1]+2*maxVal == arr[i+1] {
				arr[i] = arr[i-1] + maxVal
			} else {
				arr[i-1] = arr[i] - maxVal
			}

			break
		}

		if diff != maxVal {
			if arr[i-1]+2*maxVal == arr[i+1] {
				arr[i] = arr[i-1] + maxVal
			} else {
				arr[i-1] = arr[i] - maxVal
			}
			break
		}
	}

	fmt.Printf("output :: {%s}\n", joinInts(arr, ","))

	return arr
}
