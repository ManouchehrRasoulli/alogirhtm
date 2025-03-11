package day2

import (
	"log"
	"strconv"
	"strings"
)

func isIncreasingOrDecreasingByMaxFactor(codes []int, maxFactor int) bool {
	if len(codes) <= 1 {
		return true
	}

	asc := false
	if codes[0] < codes[1] {
		asc = true
	}

	for i := 0; i < len(codes)-1; i++ {
		a, b := codes[i], codes[i+1]
		if a == b { // nether ascending nor descending
			return false
		}

		if !asc && b > a {
			return false
		}

		if asc && a > b {
			return false
		}

		if asc {
			b, a = a, b
		}

		if a-b > maxFactor {
			return false
		}
	}

	return true
}

func isIncreasingOrDecreasingByMaxFactorByOne(codes []int, maxFactor int) bool {
	for i := range codes {
		modifiedSeq := append([]int{}, codes[:i]...)      // Copy first part
		modifiedSeq = append(modifiedSeq, codes[i+1:]...) // Add second part

		if isIncreasingOrDecreasingByMaxFactor(modifiedSeq, maxFactor) {
			// there is a valid sequence by removing one item
			return true
		}
	}

	return false
}

func ValidCodes(reader func() (string, error)) (validCodes int) {
	for {
		line, err := reader()
		if err != nil {
			return validCodes
		}

		codes := strings.Split(line, " ")
		codesInt := make([]int, 0)
		for _, v := range codes {
			i, _ := strconv.Atoi(v)
			codesInt = append(codesInt, i)
		}

		isValidCodeSequence := isIncreasingOrDecreasingByMaxFactor(codesInt, 3)
		if isValidCodeSequence {
			validCodes++
		}

		log.Println(codesInt, isValidCodeSequence)
	}
}

func ValidCodesWithTolerance(reader func() (string, error)) (validCodes int) {
	for {
		line, err := reader()
		if err != nil {
			return validCodes
		}

		codes := strings.Split(line, " ")
		codesInt := make([]int, 0)
		for _, v := range codes {
			i, _ := strconv.Atoi(v)
			codesInt = append(codesInt, i)
		}

		isValidCodeSequence := isIncreasingOrDecreasingByMaxFactorByOne(codesInt, 3)
		if isValidCodeSequence {
			validCodes++
		}

		isValid := "invalid"
		if isValidCodeSequence {
			isValid = "valid"
		}
		log.Println(codesInt, isValid)
	}
}
