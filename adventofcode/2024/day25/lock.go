package day25

import (
	"fmt"
	"log"
)

type Lock struct {
	combination []int
}

func NewLock(combination []int) *Lock {
	for i, v := range combination {
		combination[i] = v - 1
	}

	return &Lock{combination: combination}
}

func (l *Lock) TryFits(keys []*Key) int {
	// count the number of keys that fit into the lock
	var (
		total = 0
	)

	for i := 0; i < len(keys); i++ {
		keyCombination := keys[i].combination
		fits := true
		for j, v := range keyCombination {
			if l.combination[j]+v >= maxHeight {
				fits = false
				break
			}
		}

		if fits {
			log.Println(keys[i], " fits into --> ", l)
			total++
		}
	}

	return total
}

func (l *Lock) String() string {
	return fmt.Sprintf("🔒 %v", l.combination)
}
