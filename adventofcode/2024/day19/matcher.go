package day19

import "strings"

func NaiveMatch(input string, patters []string) bool {
	if len(input) == 0 {
		return true
	}

	for _, pattern := range patters {
		if strings.HasPrefix(input, pattern) {
			if NaiveMatch(input[len(pattern):], patters) {
				return true
			}
		}
	}

	return false
}

func Match(input string, patterns []string) bool {
	var (
		n = len(input)
		// use dp inorder to decrease runtime
		checkedSequence = make([]bool, n+1)
	)
	checkedSequence[0] = true // the first case means that the zero len sequence is matched

	for i := 1; i < n+1; i++ {
		checkedSequence[i] = false
		for _, pattern := range patterns {
			patternLen := len(pattern)
			if i >= patternLen && input[i-patternLen:i] == pattern {
				checkedSequence[i] = checkedSequence[i] || checkedSequence[i-patternLen]
			}
		}
	}

	return checkedSequence[n]
}

func MatchCount(input string, patterns []string) int {
	var (
		n               = len(input)
		checkedSequence = make([]int, n+1)
	)
	checkedSequence[0] = 1

	for i := 1; i < n+1; i++ {
		for _, pattern := range patterns {
			patternLen := len(pattern)
			if i >= patternLen && input[i-patternLen:i] == pattern {
				checkedSequence[i] += checkedSequence[i-patternLen]
			}
		}
	}

	return checkedSequence[n]
}
