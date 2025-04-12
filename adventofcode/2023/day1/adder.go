package day1

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Add(input string) int {
	var (
		re    = regexp.MustCompile(`\d`)
		lines = strings.Split(input, "\n")
		sum   = 0
	)

	for _, line := range lines {
		digits := re.FindAllString(line, -1)
		digitStr := digits[0] + digits[len(digits)-1]
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
			log.Fatal(err)
		}

		sum += digit
	}

	return sum
}
