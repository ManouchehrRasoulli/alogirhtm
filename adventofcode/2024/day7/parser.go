package day7

import (
	"strconv"
	"strings"
)

func ReadInRecordPart1(in string) []Record {
	lines := strings.Split(in, "\n")

	records := make([]Record, 0)
	for _, line := range lines {
		r := Record{
			FinalValue:  0,
			Numbers:     make([]int, 0),
			IsReachable: false,
		}
		lineItems := strings.Split(line, ":")

		r.FinalValue, _ = strconv.Atoi(lineItems[0])
		numbersPart := strings.TrimPrefix(lineItems[1], " ")
		numbers := strings.Split(numbersPart, " ")
		for _, number := range numbers {
			i, _ := strconv.Atoi(number)
			r.Numbers = append(r.Numbers, i)
		}
		records = append(records, r)
	}

	return records
}
