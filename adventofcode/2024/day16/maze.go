package day16

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"math"
)

const (
	start = 'S'
	end   = 'E'
	empty = '.'
	wall  = '#'
	path  = 'O'
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

type LocationCost struct {
	location *helper.Location
	cost     int
}

type State struct {
	cost          int // till this point
	loc           *helper.Location
	direction     helper.Direction
	locationCosts []LocationCost
}

func (s *State) CaptureLoscationCost(locationCosts ...LocationCost) {
	if s.locationCosts == nil {
		s.locationCosts = make([]LocationCost, 0)
	}

	s.locationCosts = append(s.locationCosts, locationCosts...)
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

func (m *Maze) PrintPath(location ...LocationCost) {
	for _, loc := range location {
		x, y := loc.location.Get()
		m.grid[x][y] = path
	}
}

func (m *Maze) PathCount() int {
	count := 0

	for _, row := range m.grid {
		for _, col := range row {
			if col == path {
				count++
			}
		}
	}

	return count
}

// SolvePart1
// return the final cost of solving the maze
func (m *Maze) SolvePart1() int {
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

// SolvePart2
// return the final cost of solving the maze
func (m *Maze) SolvePart2() (int, int) {
	var (
		visited   = make(map[Visited]struct{})
		pq        = helper.NewPriorityQueue()
		minCost   = math.MaxInt
		allPaths  = make([][]LocationCost, 0)
		finalPath []LocationCost
		paths     = 0
	)

	st := State{
		cost:      0,
		loc:       m.start,
		direction: helper.Right,
	}
	pq.Push(st)

	for pq.Len() > 0 {
		current := (*pq.Pop()).(State)
		current.CaptureLoscationCost(LocationCost{
			location: current.loc,
			cost:     current.cost,
		})

		if current.loc.Compare(m.end) == 0 {
			if current.cost < minCost {
				finalPath = current.locationCosts
				paths++
				minCost = current.cost
			}

			continue
		}

		vis := Visited{
			loc:       *current.loc,
			direction: current.direction,
		}

		if _, ok := visited[vis]; ok {
			allPaths = append(allPaths, current.locationCosts)
			continue
		}

		visited[vis] = struct{}{}

		newLoc := current.loc.Move(current.direction)
		if m.Value(newLoc) != wall {
			nstate := State{
				cost:      current.cost + 1,
				loc:       newLoc,
				direction: current.direction,
			}
			nstate.CaptureLoscationCost(current.locationCosts...)

			pq.Push(nstate)
		}

		left, right := NewDirectionLeftAndRight(current.direction)
		leftLoc := current.loc.Move(left)
		rightLoc := current.loc.Move(right)

		if m.Value(leftLoc) != wall {
			lstate := State{
				cost:      current.cost + 1001,
				loc:       leftLoc,
				direction: left,
			}
			lstate.CaptureLoscationCost(current.locationCosts...)

			pq.Push(lstate)
		}

		if m.Value(rightLoc) != wall {
			rstate := State{
				cost:      current.cost + 1001,
				loc:       rightLoc,
				direction: right,
			}
			rstate.CaptureLoscationCost(current.locationCosts...)

			pq.Push(rstate)
		}
	}

	m.PrintPath(finalPath...)
	m.RecursiveColliedWithEachOther(allPaths, [][]LocationCost{finalPath}, 3)

	return minCost, m.PathCount()
}

func (m *Maze) RecursiveColliedWithEachOther(paths [][]LocationCost, finalPaths [][]LocationCost, count int) {
	var (
		newPaths      = make([][]LocationCost, 0)
		newFinalPaths = make([][]LocationCost, 0)
	)

	newFinalPaths = append(newFinalPaths, finalPaths...)

	for _, pth := range paths {
		for _, finalPath := range finalPaths {
			if PathsColliedWithEachOther(pth, finalPath) {
				m.PrintPath(pth...)
				newFinalPaths = append(newFinalPaths, pth)
			} else {
				newPaths = append(newPaths, pth)
			}
		}
	}

	if len(newFinalPaths) > len(finalPaths) && count > 0 {
		m.RecursiveColliedWithEachOther(newPaths, newFinalPaths, count-1)
	}
}

func PathsColliedWithEachOther(path1 []LocationCost, path2 []LocationCost) bool {
	pth1 := path1[len(path1)-1]
	for _, pth := range path2 {
		if pth1.location.Compare(pth.location) == 0 && pth1.cost <= pth.cost {
			return true
		}
	}

	return false
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
