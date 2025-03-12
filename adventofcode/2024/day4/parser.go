package day4

import "strings"

func ParseInput(reader func() (string, error)) [][]string {
	data := make([][]string, 0)
	for {
		line, err := reader()
		if err != nil {
			break
		}

		lineSeparated := strings.Split(line, "")
		data = append(data, lineSeparated)
	}

	return data
}
