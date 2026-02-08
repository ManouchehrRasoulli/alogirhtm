package day11

import "strings"

const (
	start string = "you"
	end   string = "out"
)

type node struct {
	label      string
	reached    bool
	reachPaths int
	neighbors  []*node
}

func (n *node) uniquePaths() int {
	if n.label == end {
		return 1
	}

	if n.reached {
		return n.reachPaths
	}

	for _, ng := range n.neighbors {
		n.reachPaths += ng.uniquePaths()
	}

	n.reached = true
	return n.reachPaths
}

type graph struct {
	start *node
	nodes map[string]*node
}

func parseGraph(input string) *graph {
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
				label:      current,
				reachPaths: 0,
				neighbors:  make([]*node, 0),
			}
		}

		for i := 1; i < len(indicators); i++ {
			ng, ngOk := g.nodes[indicators[i]]
			if !ngOk {
				ng = &node{
					label:      indicators[i],
					reachPaths: 0,
					neighbors:  make([]*node, 0),
				}
			}
			nd.neighbors = append(nd.neighbors, ng)
			g.nodes[indicators[i]] = ng
		}

		if nd.label == start {
			g.start = nd
		}
		g.nodes[current] = nd
	}

	return &g
}
