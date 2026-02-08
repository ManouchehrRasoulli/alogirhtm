package day11

import (
	"strings"
)

const (
	// part 1
	you string = "you"

	// part 2
	start      string = "svr"
	indicator1 string = "dac"
	indicator2 string = "fft"
	end        string = "out"
)

type node struct {
	label     string
	neighbors []*node

	// part 1 cache
	reached    bool
	reachPaths int

	// part 2 cache: key = [seenDAC, seenFFT]
	memo map[[2]bool]int
}

func (n *node) uniquePathsPart1() int {
	if n.label == end {
		n.reached = true
		return 1
	}

	if n.reached {
		return n.reachPaths
	}

	for _, ng := range n.neighbors {
		n.reachPaths += ng.uniquePathsPart1()
	}

	n.reached = true
	return n.reachPaths
}

func (n *node) uniquePathsPart2(seenDAC bool, seenFFT bool) int {
	// update state based on current node
	if n.label == indicator1 {
		seenDAC = true
	}
	if n.label == indicator2 {
		seenFFT = true
	}

	key := [2]bool{seenDAC, seenFFT}
	if val, ok := n.memo[key]; ok {
		return val
	}

	// base case
	if n.label == end {
		if seenDAC && seenFFT {
			return 1
		}
		return 0
	}

	count := 0
	for _, ng := range n.neighbors {
		count += ng.uniquePathsPart2(seenDAC, seenFFT)
	}

	n.memo[key] = count
	return count
}

type graph struct {
	start *node
	nodes map[string]*node
}

func parseGraph(input string, st string) *graph {
	g := &graph{
		start: nil,
		nodes: make(map[string]*node),
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		current := strings.TrimRight(parts[0], ":")

		nd, ok := g.nodes[current]
		if !ok {
			nd = &node{
				label:     current,
				neighbors: make([]*node, 0),
				memo:      make(map[[2]bool]int),
			}
			g.nodes[current] = nd
		}

		for i := 1; i < len(parts); i++ {
			lbl := parts[i]
			ng, ok := g.nodes[lbl]
			if !ok {
				ng = &node{
					label:     lbl,
					neighbors: make([]*node, 0),
					memo:      make(map[[2]bool]int),
				}
				g.nodes[lbl] = ng
			}
			nd.neighbors = append(nd.neighbors, ng)
		}

		if current == st {
			g.start = nd
		}
	}

	return g
}
