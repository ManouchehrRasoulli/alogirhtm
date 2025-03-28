package day23

import (
	"regexp"
	"slices"
	"strings"
)

var (
	chiefConnectionPattern = regexp.MustCompile("t[a-z]")
)

type Lan struct {
	computers map[string]map[string]struct{}
}

func NewLan() *Lan {
	l := Lan{
		computers: make(map[string]map[string]struct{}),
	}

	return &l
}

func (l *Lan) MakeConnections(data []string) {
	for _, line := range data {
		parts := strings.Split(line, "-")
		cp1, cp2 := parts[0], parts[1] // connected two computer

		if _, ok := l.computers[cp1]; !ok {
			l.computers[cp1] = make(map[string]struct{})
		}

		l.computers[cp1][cp2] = struct{}{}

		if _, ok := l.computers[cp2]; !ok {
			l.computers[cp2] = make(map[string]struct{})
		}

		l.computers[cp2][cp1] = struct{}{}
	}
}

func (l *Lan) ChiefComputerTriples() int {
	var (
		uniqueTriples = make(map[string]struct{})
	)

	for cp1, _ := range l.computers {
		for cp2, _ := range l.computers[cp1] {
			for cp3, _ := range l.computers[cp2] {
				triples := []string{cp1, cp2, cp3}
				if _, ok := l.computers[cp3][cp1]; !ok {
					continue
				}

				slices.Sort(triples)
				triplesStr := strings.Join(triples, ",")
				if chiefConnectionPattern.MatchString(triplesStr) {
					uniqueTriples[triplesStr] = struct{}{}
				}
			}
		}
	}

	return len(uniqueTriples)
}
