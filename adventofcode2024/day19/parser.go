package day19

import (
	"strings"
)

func Parse(input string) ([]Replacement, string) {
	lines := strings.Split(input, "\n")
	var (
		replacements = make([]Replacement, 0)
		data         string
	)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, "=>") {
			parts := strings.Split(line, " => ")
			replacements = append(replacements, Replacement{
				X:    parts[0],
				With: parts[1],
			})
		} else {
			data = strings.Trim(line, " ")
		}
	}

	return replacements, data
}

func Replacements(replacements []Replacement, data string) []string {
	with := make([]string, 0)
	for _, replacement := range replacements {
		if replacement.X == data {
			with = append(with, replacement.With)
		}
	}

	return with
}
