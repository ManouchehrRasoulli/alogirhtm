package _kyu

import "strings"

func splitBySize(input string, size int) []string {
	var sections []string
	var stop int
	for i := 0; i < len(input); i += size {
		stop = i + size
		if stop > len(input) {
			stop = len(input)
		}
		sections = append(sections, input[i:stop])
	}
	return sections
}

func parseKey(k string) map[rune]rune {
	kv := splitBySize(strings.ToLower(k), 2)

	m := make(map[rune]rune)
	for _, i := range kv {
		m[rune(i[0])] = rune(i[1])
		m[rune(i[1])] = rune(i[0])
		v := strings.ToUpper(i)
		m[rune(v[0])] = rune(v[1])
		m[rune(v[1])] = rune(v[0])
	}

	return m
}

func Encode(str, key string) string {
	km := parseKey(key)

	return strings.Map(func(r rune) rune {
		if v, ok := km[r]; ok {
			return v
		}
		return r
	}, str)
}

func Decode(str, key string) string {
	km := parseKey(key)

	return strings.Map(func(r rune) rune {
		if v, ok := km[r]; ok {
			return v
		}
		return r
	}, str)
}
