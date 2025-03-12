package day5

import "fmt"

func Interpret(precedenceMap PrecedenceHierarchy, updates []Update) {
	for ui, update := range updates {
		var (
			isValid = true
			cause   = make([]string, 0)
		)
		for i := 0; i < len(update.Items); i++ {
			for j := i + 1; j < len(update.Items); j++ {
				if precedenceMap.Precede(update.Items[i], update.Items[j]) {
					isValid = false
					cause = append(cause, fmt.Sprintf("%d: %d proceede %d: %d", j, update.Items[j], i, update.Items[i]))
				}
			}
		}
		if !isValid {
			updates[ui].IsValid = false
			updates[ui].Cause = cause
		}
	}
}

func ValidSum(updates []Update) int {
	sum := 0
	for _, update := range updates {
		if update.IsValid {
			sum += update.Items[len(update.Items)/2]
		}
	}
	return sum
}
