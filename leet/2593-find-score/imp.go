package _593_find_score

import (
	"sort"
)

type element struct {
	value int
	index int
}

func findScore(nums []int) int64 {
	n := len(nums)

	elements := make([]element, n)
	for i, num := range nums {
		elements[i] = element{num, i}
	}

	sort.Slice(elements, func(i, j int) bool {
		if elements[i].value == elements[j].value {
			return elements[i].index < elements[j].index
		}

		return elements[i].value < elements[j].value
	})

	marked := make([]bool, n)
	score := int64(0)

	for _, element := range elements {
		if !marked[element.index] {
			score += int64(element.value)
			marked[element.index] = true
			if element.index > 0 {
				marked[element.index-1] = true
			}
			if element.index < n-1 {
				marked[element.index+1] = true
			}
		}
	}

	return score
}
