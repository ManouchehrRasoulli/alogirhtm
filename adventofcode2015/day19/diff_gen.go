package day19

import (
	"hash/fnv"
	"log"
	"strings"
)

func HashFNV1a(input string) uint32 {
	hsh := fnv.New32a()
	_, _ = hsh.Write([]byte(input))
	return hsh.Sum32()
}

func FindSubstringFromIndex(input, sub string, startIndex int) int {
	if startIndex >= len(input) {
		return -1 // Out of range
	}
	slice := input[startIndex:]
	index := strings.Index(slice, sub)

	if index == -1 {
		return -1
	}
	return startIndex + index
}

func ReplaceSubstringAt(input, old, new string, startIndex int) string {
	pos := strings.Index(input[startIndex:], old)
	if pos == -1 {
		return input
	}
	pos += startIndex
	return input[:pos] + new + input[pos+len(old):]
}

func Diffs(replacements []Replacement, data string) int {
	saw := make(map[uint32]struct{})
	for _, replacement := range replacements {
		replacementIndex := FindSubstringFromIndex(data, replacement.X, 0)
		for replacementIndex != -1 {
			newStr := ReplaceSubstringAt(data, replacement.X, replacement.With, replacementIndex)
			saw[HashFNV1a(newStr)] = struct{}{}
			replacementIndex = FindSubstringFromIndex(data, replacement.X, replacementIndex+1)
		}
	}

	log.Printf("%+v\n", saw)

	return len(saw)
}

func Reduce(replacements []Replacement, data, target string) int {
	reverseMap := make(map[string]string)
	for _, replacement := range replacements {
		reverseMap[replacement.With] = replacement.X
	}

	steps := 0

	log.Printf("%+v\n", reverseMap)
	for data != target {
		replaced := false
		for to, from := range reverseMap {
			if strings.Contains(data, to) {
				data = strings.Replace(data, to, from, 1)
				steps++
				replaced = true
				break
			}
		}

		if !replaced {
			break
		}
	}

	return steps
}
