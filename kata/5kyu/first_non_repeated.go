package _kyu

import (
	"strings"
)

func FirstNonRepeating(str string) string {
	charMap := make(map[string]int)

	for i := 0; i < len(str); i++ {
		charMap[strings.ToLower(string(str[i]))]++
	}

	for i := 0; i < len(str); i++ {
		if charMap[strings.ToLower(string(str[i]))] == 1 {
			return string(str[i])
		}
	}

	return ""
}
