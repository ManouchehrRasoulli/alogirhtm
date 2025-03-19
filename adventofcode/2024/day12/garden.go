package day12

import (
	"container/list"
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"log"
)

const (
	noValue = '&'
)

var (
	directions = []helper.Direction{
		helper.Left,
		helper.Right,
		helper.Top,
		helper.Bottom,
	}
)

type Location struct {
	x int
	y int
}

func (l *Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.x, l.y)
}

type Edge struct {
	a, b Location
}

func NormalizeEdge(a, b Location) Edge {
	if a.x < b.x || (a.x == b.x && a.y < b.y) {
		return Edge{a, b}
	}
	return Edge{b, a}
}

type Garden struct {
	Plots   [][]rune
	Visited map[Location]struct{}
}

func (g *Garden) Value(location Location) rune {
	if location.x < 0 || location.y < 0 || location.x >= len(g.Plots) || location.y >= len(g.Plots) {
		return noValue
	}

	return g.Plots[location.x][location.y]
}

func (g *Garden) FencesAmountPart1() int {
	amount := 0

	for i := 0; i < len(g.Plots); i++ {
		for j := 0; j < len(g.Plots[i]); j++ {
			grid, perimeters := g.Flood(Location{i, j})
			amount += grid * perimeters
		}
	}

	return amount
}

// Flood
// this implementation for part 1 will flood a location with bfs search and then
// return the number of grid and perimeters
func (g *Garden) Flood(location Location) (int, int) {
	var (
		grids     = 0
		perimeter = 0
		queue     = list.New()
	)

	queue.PushBack(location)
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Location)
		if _, visited := g.Visited[current]; visited {
			continue
		}

		g.Visited[current] = struct{}{}
		ngs := g.Neighbours(current)
		grids += 1

		currentPerimeter := 4
		for _, ng := range ngs {
			queue.PushBack(ng)
			currentPerimeter--
		}
		perimeter += currentPerimeter
	}

	return grids, perimeter
}

func (g *Garden) FencesAmountPart2() int {
	amount := 0

	for i := 0; i < len(g.Plots); i++ {
		for j := 0; j < len(g.Plots[i]); j++ {
			loc := Location{i, j}
			grid, edges := g.FloodGetEdges(loc)
			if grid > 0 {
				log.Println(string(g.Value(loc)), grid, edges)
			}
			amount += grid * edges
		}
	}

	return amount
}

func (g *Garden) FloodGetEdges(location Location) (int, int) {
	var (
		grids   = 0
		corners = 0
		queue   = list.New()
	)

	queue.PushBack(location)
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Location)
		if _, visited := g.Visited[current]; visited {
			continue
		}

		g.Visited[current] = struct{}{}
		grids += 1
		v := g.Value(current)

		ngs := g.Neighbours(current)
		for _, ng := range ngs {
			if g.Value(ng) == v {
				queue.PushBack(ng)
			}
		}

		ngs = g.AllNeighbours(current)
		vtl := g.Value(ngs[0])
		vt := g.Value(ngs[1])
		vtr := g.Value(ngs[2])
		vl := g.Value(ngs[3])
		vr := g.Value(ngs[4])
		vbl := g.Value(ngs[5])
		vb := g.Value(ngs[6])
		vbr := g.Value(ngs[7])

		if vt != v && vl != v {
			corners++
		}
		if vt == v && vl == v && vtl != v {
			corners++
		}
		if vt != v && vr != v {
			corners++
		}
		if vt == v && vr == v && vtr != v {
			corners++
		}
		if vb != v && vl != v {
			corners++
		}
		if vb == v && vl == v && vbl != v {
			corners++
		}
		if vb != v && vr != v {
			corners++
		}
		if vb == v && vr == v && vbr != v {
			corners++
		}
	}

	return grids, corners // each cornet contribute in two edge line
}

func (g *Garden) Neighbours(location Location) []Location {
	lv := g.Value(location)
	if lv == noValue {
		return nil
	}

	ngs := make([]Location, 0)
	for _, d := range directions {
		ng := Location{
			x: location.x + d.X(),
			y: location.y + d.Y(),
		}

		if g.Value(ng) == lv {
			ngs = append(ngs, ng)
		}
	}

	return ngs
}

func (g *Garden) AllNeighbours(location Location) []Location {
	ngs := make([]Location, 0)

	for _, d := range helper.AllDirections {
		ng := Location{
			x: location.x + d.X(),
			y: location.y + d.Y(),
		}
		ngs = append(ngs, ng)
	}

	return ngs
}

func (g *Garden) String() string {
	lines := "\n"

	for _, row := range g.Plots {
		lines += string(row) + "\n"
	}

	return lines
}
