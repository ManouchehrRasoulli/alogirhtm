package regular_expression

import "fmt"

const (
	start  string = "^"
	end    string = "$"
	single string = "."
	ani    string = "*"
)

type state struct {
	finalState bool // indicate which this state is a valid state or what
	match      string
	next       *state
	prev       *state
}

func (s *state) matchString(st string, index int) bool {
	if s.match == start && index == 0 {
		return s.next.matchString(st, index)
	} else if s.match == start && index != 0 {
		return false
	}

	if s.match == end && len(st) != index {
		return false
	}

	if len(st) == index {
		return s.finalState
	}

	if s.match == single {
		return s.next.matchString(st, index+1)
	}

	if s.match == ani {
		if s.next.matchString(st, index+1) {
			return s.next.matchString(st, index+1)
		} else {
			return s.matchString(st, index+1)
		}
	}

	if s.match == string(st[index]) {
		return s.next.matchString(st, index+1)
	} else {
		return s.prev.matchString(st, index+1)
	}
}

func (s *state) String() string {
	if s.next != nil {
		return fmt.Sprintf("{f: %t, m: %s, n: %s}", s.finalState, s.match, s.next.String())
	}
	return fmt.Sprintf("{f: %t, m: %s}", s.finalState, s.match)
}

func parse(p string, currentNode *state, index int) {
	s := state{
		finalState: len(p) == index,
		match:      end,
		next:       nil,
		prev:       currentNode,
	}
	currentNode.next = &s

	if s.finalState {
		return
	}

	current := p[index]
	switch current {
	case '.':
		{
			s.match = single
			parse(p, &s, index+1)
			return
		}
	case '*':
		{
			s.match = ani
			parse(p, &s, index+1)
			return
		}
	default:
		{
			s.match = string(p[index])
			parse(p, &s, index+1)
			return
		}
	}
}

func NaiveMatch(st string, p string) bool {
	startState := state{
		finalState: false,
		match:      start,
		next:       nil,
		prev:       nil,
	}

	parse(p, &startState, 0)
	fmt.Println(startState.String())
	return startState.matchString(st, 0)
}
