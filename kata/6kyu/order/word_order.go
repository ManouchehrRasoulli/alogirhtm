package kata

import (
	"regexp"
	"strconv"
	"strings"
)

func Order(sentence string) string {
	words := strings.Split(sentence, " ")

	re := regexp.MustCompile(`\d`)
	n := len(words)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ad, bd := re.FindString(words[i]), re.FindString(words[j])

			ai, _ := strconv.Atoi(ad)
			bi, _ := strconv.Atoi(bd)

			if ai > bi {
				words[i], words[j] = words[j], words[i]
			}
		}
	}

	return strings.Join(words, " ")
}
