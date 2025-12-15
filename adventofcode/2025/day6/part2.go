package day6

import (
	"strconv"
	"strings"
)

func parseOperators(line string) ([]string, []int) {
	var (
		operators        []string
		lens             []int
		currentOperator  string
		currentColumnLen int
	)
	for _, c := range line {
		if c != ' ' {
			if currentColumnLen != 0 {
				operators = append(operators, currentOperator)
				lens = append(lens, currentColumnLen-1)
			}

			currentOperator = string(c)
			currentColumnLen = 1
			continue
		}
		currentColumnLen++
	}
	if currentColumnLen != 0 {
		operators = append(operators, currentOperator)
		lens = append(lens, currentColumnLen)
	}
	return operators, lens
}

func parseNumStrings(line string, lens []int) []string {
	var (
		nums  []string
		index = 0
	)

	for i := 0; i < len(lens); i++ {
		nums = append(nums, line[index:index+lens[i]])
		index += lens[i] + 1
	}

	return nums
}

func vertical(nums [][]string, column int) []string {
	nms := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		nms = append(nms, nums[i][column])
	}
	return nms
}

func verticalOp(nums []string, op string, ln int) int {
	totalValue := 0
	if op == "*" {
		totalValue = 1
	}

	for i := 0; i < ln; i++ {
		numStr := strings.Builder{}
		for j := 0; j < len(nums); j++ {
			_, _ = numStr.WriteString(string(nums[j][i]))
		}
		numVal, _ := strconv.Atoi(strings.TrimSpace(numStr.String()))
		switch op {
		case "*":
			totalValue *= numVal
		case "+":
			totalValue += numVal
		}
	}

	return totalValue
}

func somUp(nums [][]string, operators []string, ln []int) int {
	total := 0
	for i, op := range operators {
		nms := vertical(nums, i)
		total += verticalOp(nms, op, ln[i])
	}
	return total
}

func SolveHorizontallyAndSum(input string) int {
	var (
		lines     = strings.Split(input, "\n")
		operators []string
		lens      []int
		nums      [][]string
	)

	operators, lens = parseOperators(lines[len(lines)-1])
	for i := 0; i < len(lines)-1; i++ {
		nums = append(nums, parseNumStrings(lines[i], lens))
	}

	return somUp(nums, operators, lens)
}
