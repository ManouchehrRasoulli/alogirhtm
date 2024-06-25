package _kyu

import "sort"

func TwoSum(numbers []int, target int) [2]int {
	var (
		i = 0
		j = len(numbers) - 1
	)

	sort.Ints(numbers)

loop:
	for i != j {
		if numbers[j] > target {
			j--
			continue
		}
		switch t := numbers[i] + numbers[j]; {
		case t > target:
			j--
			continue
		case t < target:
			i++
		default:
			break loop
		}

	}

	if i == j { // didn't find the items
		return [2]int{}
	}

	return [2]int{i, j}
}
