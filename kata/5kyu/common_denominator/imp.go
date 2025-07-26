package kata

import (
	"fmt"
	"strings"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmMultiple(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}

	return result
}

func simplify(n, d int) (int, int) {
	g := gcd(n, d)
	return n / g, d / g
}

func ConvertFracts(a [][]int) string {
	for i := range a {
		a[i][0], a[i][1] = simplify(a[i][0], a[i][1])
	}

	denominators := make([]int, len(a))
	for i, frac := range a {
		denominators[i] = frac[1]
	}

	commonDenom := lcmMultiple(denominators)

	var b strings.Builder
	for _, frac := range a {
		num := frac[0]
		den := frac[1]
		adjustedNum := num * (commonDenom / den)
		b.WriteString(fmt.Sprintf("(%d,%d)", adjustedNum, commonDenom))
	}

	return b.String()
}
