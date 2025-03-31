package day24

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	and = "AND"
	or  = "OR"
	xor = "XOR"
)

var (
	wirePattern        = regexp.MustCompile(`([a-z0-9]{3}): ([0-1])`)
	instructionPattern = regexp.MustCompile(`([a-z0-9]{3})\s+(AND|OR|XOR)\s+([a-z0-9]{3})\s+->\s+([a-z0-9]{3})`)
)

func DoOperations(lines []string) int {
	var (
		wires      = make(map[string]int)
		operations = make([]string, 0)
		wireParts  = true
		zWires     = make([]string, 0)
		zValues    = make([]string, 0)
	)

	for _, line := range lines {
		if wireParts {
			m := wirePattern.FindAllStringSubmatch(line, -1)
			if len(m) > 0 {
				wire := m[0][1]
				if m[0][2] == "1" {
					wires[wire] = 1
				} else {
					wires[wire] = 0
				}
				continue
			} else {
				wireParts = false
				continue
			}
		}

		operations = append(operations, line)
	}

	for len(operations) > 0 {
		operation := operations[0]
		m := instructionPattern.FindAllStringSubmatch(operation, -1)
		if len(m) == 0 {
			log.Fatalf("invalid operation %q", operation)
		}

		ia := m[0][1]
		op := m[0][2]
		ib := m[0][3]
		o := m[0][4]

		aw, af := wires[ia]
		bw, bf := wires[ib]

		if !af || !bf { // inputs didn't have a valid value yet !
			operations = operations[1:]
			operations = append(operations, operation)
			continue
		}

		switch op {
		case and:
			wires[o] = aw & bw
		case or:
			wires[o] = aw | bw
		case xor:
			wires[o] = aw ^ bw
		}

		// operation completed
		operations = operations[1:]
	}

	for k := range wires {
		if strings.HasPrefix(k, "z") {
			zWires = append(zWires, k)
		}
	}

	slices.Sort(zWires)
	slices.Reverse(zWires)
	for _, v := range zWires {
		zValues = append(zValues, fmt.Sprintf("%d", wires[v]))
	}

	zString := strings.Join(zValues, "")
	zInt, err := strconv.ParseInt(zString, 2, 64)
	if err != nil {
		log.Fatal(err, zString)
	}

	return int(zInt)
}

func FixCircuit(lines []string) string {
	// have to implement a backtracking method to modify circuit and test if
	// the output is correct or not
	panic("not implemented")
}
