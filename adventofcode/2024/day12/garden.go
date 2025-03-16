package day12

import (
	"container/list"
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/day4"
)

const (
	noValue = '.'
)

var (
	directions = []day4.Direction{
		day4.Top,
		day4.Left,
		day4.Right,
		day4.Bottom,
	}
)

type Location struct {
	x int
	y int
}

func (l *Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.x, l.y)
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

func (g *Garden) String() string {
	lines := "\n"

	for _, row := range g.Plots {
		lines += string(row) + "\n"
	}

	return lines
}
