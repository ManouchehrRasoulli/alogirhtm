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

func findOverlapping(text string, pattern *regexp.Regexp) []string {
	matches := make([]string, 0)
	for i := 0; i < len(text); i++ {
		loc := pattern.FindStringIndex(text[i:])
		if loc != nil && loc[0] == 0 {
			matches = append(matches, text[i:i+loc[1]])
		}
	}
	return matches
}

func AddWithStr(input string) int {
	var (
		re          = regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine|\d)`)
		lines       = strings.Split(input, "\n")
		sum         = 0
		strToDigits = map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
			"1":     1,
			"2":     2,
			"3":     3,
			"4":     4,
			"5":     5,
			"6":     6,
			"7":     7,
			"8":     8,
			"9":     9,
		}
	)

	for _, line := range lines {
		digits := findOverlapping(line, re)
		log.Println(line, digits, 10*strToDigits[digits[0]]+strToDigits[digits[len(digits)-1]])
		sum += 10*strToDigits[digits[0]] + strToDigits[digits[len(digits)-1]]
	}

	return sum
}
