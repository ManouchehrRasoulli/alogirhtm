package day12

import (
	"fmt"
)

var (
	memo map[memoKey]int
)

func Part2(lines []dataLines) int {
	var permutations int

	for _, line := range lines {
		line = line.Unfold(5)
		fmt.Println(line)
		memo = make(map[memoKey]int)
		permutations += countValid(string(line.Data), line.Match)
	}

	return permutations
}

func countValid(pattern string, groups []int) int {
	key := memoKey{pattern, fmt.Sprint(groups)}
	if val, ok := memo[key]; ok {
		return val
	}

	if len(groups) == 0 {
		for _, c := range pattern {
			if c == '#' {
				memo[key] = 0
				return 0
			}
		}
		memo[key] = 1
		return 1
	}

	total := 0
	groupLen := groups[0]

	for i := 0; i <= len(pattern)-groupLen; i++ {
		segment := pattern[i : i+groupLen]
		if isValidSegment(segment) {
			if i+groupLen == len(pattern) || pattern[i+groupLen] != '#' {
				rest := ""
				if i+groupLen+1 < len(pattern) {
					rest = pattern[i+groupLen+1:]
				}
				total += countValid(rest, groups[1:])
			}
		}
		if pattern[i] == '#' {
			break
		}
	}

	memo[key] = total
	return total
}

func isValidSegment(segment string) bool {
	for _, c := range segment {
		if c != '#' && c != '?' {
			return false
		}
	}
	return true
}
