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
				if _, ok := l.computers[cp3][cp1]; !ok {
					continue
				}

				triples := []string{cp1, cp2, cp3}
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

func (l *Lan) LargestClique() string {
	var (
		cliques   = l.cliques()
		maxClique = make([]string, 0)
	)

	for _, clique := range cliques {
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}

	slices.Sort(maxClique)
	return strings.Join(maxClique, ",")
}

func (l *Lan) cliques() [][]string {
	var (
		maxClique  = make([][]string, 0)
		candidates = make([]string, 0)
	)

	for k, _ := range l.computers {
		candidates = append(candidates, k)
	}

	dfs(l.computers, []string{}, candidates, &maxClique)

	return maxClique
}

func dfs(graph map[string]map[string]struct{}, currentClique []string, candidates []string, maxClique *[][]string) {
	if len(candidates) == 0 {
		*maxClique = append(*maxClique, currentClique)
		return
	}

	for i, v := range candidates {
		var (
			newClique     = append(currentClique, v)
			newCandidates = make([]string, 0)
		)

		for _, c := range candidates[i+1:] {
			if _, ok := graph[v][c]; ok {
				newCandidates = append(newCandidates, c)
			}
		}

		dfs(graph, newClique, newCandidates, maxClique)
	}
}
