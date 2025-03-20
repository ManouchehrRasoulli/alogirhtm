package day16

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
)

const (
	start = 'S'
	end   = 'E'
	empty = '.'
	wall  = '#'
)

var (
	direction = []helper.Direction{
		helper.Top,
		helper.Right,
		helper.Bottom,
		helper.Left,
	}
)

func NewDirectionLeftAndRight(d helper.Direction) (left helper.Direction, right helper.Direction) {
	for i, v := range direction {
		if v == d {
			il := (i + 3) % 4
			left = direction[il]

			ir := (i + 1) % 4
			right = direction[ir]

			break
		}
	}

	return left, right
}

type State struct {
	cost      int // till this point
	loc       *helper.Location
	direction helper.Direction
}

func (s State) Value() int {
	return s.cost // priority is based on cost of the position
}

var _ helper.PriorityQueueElement = State{}

type Visited struct {
	loc       helper.Location
	direction helper.Direction
}

type Maze struct {
	grid  [][]rune
	start *helper.Location
	end   *helper.Location
}

func (m *Maze) Value(loc *helper.Location) rune {
	if loc == nil {
		return wall
	}

	x, y := loc.Get()
	if x >= 0 && y >= 0 && x < len(m.grid) && y < len(m.grid[x]) {
		return m.grid[x][y]
	}

	return wall
}

// Solve
// return the final cost of solving the maze
func (m *Maze) Solve() int {
	var (
		visited = make(map[Visited]struct{})
		pq      = helper.NewPriorityQueue()
	)

	st := State{
		cost:      0,
		loc:       m.start,
		direction: helper.Right,
	}
	pq.Push(st)

	for pq.Len() > 0 {
		current := (*pq.Pop()).(State)
		if current.loc.Compare(m.end) == 0 {
			return current.cost
		}

		vis := Visited{
			loc:       *current.loc,
			direction: current.direction,
		}

		if _, ok := visited[vis]; ok {
			continue
		}

		visited[vis] = struct{}{}

		newLoc := current.loc.Move(current.direction)
		if m.Value(newLoc) != wall {
			pq.Push(State{
				cost:      current.cost + 1,
				loc:       newLoc,
				direction: current.direction,
			})
		}

		left, right := NewDirectionLeftAndRight(current.direction)
		leftLoc := current.loc.Move(left)
		rightLoc := current.loc.Move(right)

		if m.Value(leftLoc) != wall {
			pq.Push(State{
				cost:      current.cost + 1001,
				loc:       leftLoc,
				direction: left,
			})
		}

		if m.Value(rightLoc) != wall {
			pq.Push(State{
				cost:      current.cost + 1001,
				loc:       rightLoc,
				direction: right,
			})
		}
	}

	return -1 // no path found
}

func (m *Maze) String() string {
	data := "\n🐍\n"

	for _, row := range m.grid {
		for _, char := range row {
			data += string(char)
		}
		data += "\n"
	}

	if m.start != nil {
		data += fmt.Sprintf("🤖 %s\n", m.start.String())
	}

	if m.end != nil {
		data += fmt.Sprintf("🔚 %s\n", m.end.String())
	}

	return data
}
