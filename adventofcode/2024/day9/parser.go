package day9

import (
	"strings"
)

func ReadInput(data string) []rune {
	lines := strings.Split(data, "\n")

	runes := make([]rune, 0)
	for _, line := range lines {
		runes = append(runes, []rune(line)...)
	}

	return runes
}
