package day5

import "slices"

func Interpret(precedenceMap PrecedenceHierarchy, updates []Update) {
	for ui, update := range updates {
		var (
			isValid = true
			cause   = make([]Cause, 0)
		)
		for i := 0; i < len(update.Items); i++ {
			for j := i + 1; j < len(update.Items); j++ {
				if precedenceMap.Precede(update.Items[i], update.Items[j]) {
					isValid = false
					cause = append(cause, Cause{
						JIndex: j, JValue: update.Items[j], IIndex: i, IValue: update.Items[i],
					})
				}
			}
		}
		if !isValid {
			updates[ui].IsValid = false
			updates[ui].Cause = cause
		} else {
			updates[ui].IsValid = true
			updates[ui].Cause = make([]Cause, 0)
		}
	}
}

func SortFunc(precedence PrecedenceHierarchy) func(i, j int) int {
	return func(i, j int) int {
		if precedence.Precede(j, i) {
			return -1
		}

		if precedence.Precede(i, j) {
			return 1
		}

		return 0
	}
}

func Fix(updates []Update, precedence PrecedenceHierarchy) {
	sortFunc := SortFunc(precedence)

	for ui, update := range updates {
		if !update.IsValid {
			slices.SortFunc(updates[ui].Items, sortFunc)
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
