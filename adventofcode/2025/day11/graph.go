package day11

import "strings"

const (
	// part 1
	you string = "you"

	// part 2
	start      string = "svr"
	indicator1 string = "dac"
	indicator2 string = "fft"

	end string = "out"
)

type node struct {
	label     string
	neighbors []*node
	reached   bool

	// only for part 1
	reachPaths int

	// only for part 2

}

func (n *node) uniquePathsPart1() int {
	if n.label == end {
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

type graph struct {
	start *node
	nodes map[string]*node
}

func parseGraphPart1(input string) *graph {
	var (
		g = graph{
			start: nil,
			nodes: make(map[string]*node),
		}
		lines = strings.Split(input, "\n")
	)

	for _, line := range lines {
		indicators := strings.Split(line, " ")
		current := strings.TrimRight(indicators[0], ":")
		nd, ok := g.nodes[current]
		if !ok {
			nd = &node{
				label:     current,
				neighbors: make([]*node, 0),
			}
		}

		for i := 1; i < len(indicators); i++ {
			ng, ngOk := g.nodes[indicators[i]]
			if !ngOk {
				ng = &node{
					label:     indicators[i],
					neighbors: make([]*node, 0),
				}
			}
			nd.neighbors = append(nd.neighbors, ng)
			g.nodes[indicators[i]] = ng
		}

		if nd.label == you {
			g.start = nd
		}
		g.nodes[current] = nd
	}

	return &g
}
