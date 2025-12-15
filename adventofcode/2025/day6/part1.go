package day6

import (
	"log"
	"strconv"
	"strings"
)

func SolveAndSum(input string) int {
	var (
		lines      = strings.Split(input, "\n")
		numbers    = make([][]int, 0)
		operations = make([]string, 0)
		totalValue = 0
		sumFunc    = func(nums []int) int {
			total := 0
			for _, n := range nums {
				total += n
			}
			return total
		}
		multiFunc = func(nums []int) int {
			total := 1
			if len(nums) == 0 {
				return 0
			}
			for _, n := range nums {
				total *= n
			}
			return total
		}
		subFunc = func(nums []int) int {
			total := 0
			for _, n := range nums {
				total -= n
			}
			return total
		}
	)

	for i, line := range lines {
		var (
			nums  = make([]int, 0)
			items = strings.Split(line, " ")
		)

		for _, item := range items {
			if len(item) == 0 {
				continue
			}

			if i == len(lines)-1 {
				operations = append(operations, item)
				continue
			}

			num, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err, line)
			}
			nums = append(nums, num)
		}
		if len(nums) == 0 {
			continue
		}

		numbers = append(numbers, nums)
	}

	for i := 0; i < len(operations); i++ {
		var (
			nums      = make([]int, 0)
			operation = operations[i]
		)
		for j := 0; j < len(numbers); j++ {
			nums = append(nums, numbers[j][i])
		}

		switch operation {
		case "+":
			totalValue += sumFunc(nums)
		case "*":
			totalValue += multiFunc(nums)
		case "-":
			totalValue += subFunc(nums)
		}
	}

	return totalValue
}
