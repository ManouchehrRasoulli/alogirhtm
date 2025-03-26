package day21

import (
	"log"
	"regexp"
	"strconv"
)

var (
	numberRegex = regexp.MustCompile(`[0-9]+`)
)

func ComputeValues(sequence string, finalPath []rune) int {
	numStr := numberRegex.FindAllString(sequence, 1)[0]
	num, _ := strconv.Atoi(numStr)

	log.Println("----", num, sequence, len(finalPath))

	return num * (len(finalPath))
}
