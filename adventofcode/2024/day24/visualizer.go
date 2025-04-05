package day24

import (
	"fmt"
	"log"
	"strings"
)

type Edge struct {
	Label string
	Value int
}

func (e *Edge) String() string {
	return fmt.Sprintf("%s [label=\"%s -> %b\", shape=ellipse];\n", e.Label, e.Label, e.Value)
}

type Node struct {
	Label  string
	InputA string
	InputB string
	Output string
}

func (node *Node) String() string {
	var data string

	data += "\n"

	// define the logical gate
	if strings.Contains(node.Label, and) {
		data += fmt.Sprintf("%s [label=\"AND\", shape=box];\n", node.Label)
	} else if strings.Contains(node.Label, xor) {
		data += fmt.Sprintf("%s [label=\"XOR\", shape=box];\n", node.Label)
	} else if strings.Contains(node.Label, or) {
		data += fmt.Sprintf("%s [label=\"OR\", shape=box];\n", node.Label)
	} else {
		log.Fatalf("invalid operation %q", node.Label)
	}

	data += fmt.Sprintf("%s -> %s;\n", node.InputA, node.Label)
	data += fmt.Sprintf("%s -> %s;\n", node.InputB, node.Label)
	data += fmt.Sprintf("%s -> %s;\n", node.Label, node.Output)

	return data
}

func ToDotFormat(lines []string) string {
	var (
		edges     = make(map[string]Edge)
		nodes     = make([]Node, 0)
		wireParts = true
		output    string
	)

	output += "digraph LogicCircuit {\n"

	for i, line := range lines {
		if wireParts {
			m := wirePattern.FindAllStringSubmatch(line, -1)
			if len(m) > 0 {
				wire := m[0][1]
				if m[0][2] == "1" {
					edges[wire] = Edge{
						Label: wire,
						Value: 1,
					}
				} else {
					edges[wire] = Edge{
						Label: wire,
						Value: 0,
					}
				}
				continue
			} else { // jump one line ahead
				wireParts = false
				continue
			}
		}

		// continue with operators
		m := instructionPattern.FindAllStringSubmatch(line, -1)
		if len(m) == 0 {
			log.Fatalf("invalid operation %q", line)
		}

		ia := m[0][1]
		op := m[0][2]
		ib := m[0][3]
		o := m[0][4]

		node := Node{
			Label:  fmt.Sprintf("%s%d", op, i),
			InputA: ia,
			InputB: ib,
			Output: o,
		}

		nodes = append(nodes, node)
	}

	// sort and print the nodes
	for len(nodes) > 0 {
		node := nodes[0]

		aw, af := edges[node.InputA]
		bw, bf := edges[node.InputB]
		var out Edge

		if !af || !bf { // kind of sort the nodes
			nodes = nodes[1:]
			nodes = append(nodes, node)
			continue
		}

		if strings.Contains(node.Label, and) {
			out = Edge{
				Label: node.Output,
				Value: aw.Value & bw.Value,
			}
			edges[node.Output] = out
		} else if strings.Contains(node.Label, xor) {
			out = Edge{
				Label: node.Output,
				Value: aw.Value ^ bw.Value,
			}
			edges[node.Output] = out
		} else if strings.Contains(node.Label, or) {
			out = Edge{
				Label: node.Output,
				Value: aw.Value | bw.Value,
			}
			edges[node.Output] = out
		}

		output += aw.String()
		output += bw.String()
		output += out.String()
		output += node.String()

		// operation completed
		nodes = nodes[1:]
	}

	output += "}\n"

	return output
}
