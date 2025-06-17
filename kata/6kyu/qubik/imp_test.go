package kata

import (
	"fmt"
	"testing"
)

func TestSolomonsQuest(t *testing.T) {
	mapExample := [][3]int{
		{1, 3, 5},
		{2, 0, 10},
		{-3, 1, 4},
		{4, 2, 4},
		{1, 1, 5},
		{-3, 0, 12},
		{2, 1, 12},
		{-2, 2, 6},
	}

	result := SolomonsQuest(mapExample)
	fmt.Println(result) // Output: [346 40]
}
