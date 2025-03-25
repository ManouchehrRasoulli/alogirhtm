package day20

import (
	"container/list"
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"log"
)

const (
	freeSpace     = '.'
	usedSpace     = '#'
	trialSpace    = '*' // path which we found
	startLocation = 'S'
	endLocation   = 'E'
)

type locationPath struct {
	locations []*helper.Location
	current   *helper.Location
}

type Puzzle struct {
	grid  [][]rune
	row   int
	col   int
	start *helper.Location
	end   *helper.Location
	path  []*helper.Location
}

func (p *Puzzle) Value(x, y int) rune {
	if x < 0 || y < 0 || x >= len(p.grid) || y >= len(p.grid[x]) {
		return usedSpace
	}

	return p.grid[x][y]
}

func (p *Puzzle) WriteValue(x, y int, value rune) {
	switch p.Value(x, y) {
	case freeSpace:
		break
	case usedSpace:
		log.Fatal("Used space marked !!!", x, y, p.Value(x, y))
	case startLocation, endLocation:
		return
	}

	p.grid[x][y] = value
}

func (p *Puzzle) MarkPath(path []*helper.Location) {
	for _, l := range path {
		x, y := l.Get()
		p.WriteValue(x, y, trialSpace)
	}
}

func (p *Puzzle) BfsWithPath() bool {
	var (
		visited = make(map[helper.Location]struct{})
		queue   = list.New()
		end     = p.end
		path    = locationPath{
			locations: make([]*helper.Location, 0),
			current:   p.start,
		}
		found = false
	)

	queue.PushBack(path)
	for queue.Len() > 0 {
		currentPath := queue.Remove(queue.Front()).(locationPath)
		currentPath.locations = append(currentPath.locations, currentPath.current)
		visited[*currentPath.current] = struct{}{}

		if currentPath.current.Compare(end) == 0 {
			found = true
			p.MarkPath(currentPath.locations)
			break
		}

		if p.Value(currentPath.current.Get()) == usedSpace {
			continue
		}

		for _, dir := range helper.PathDirections {
			ng := currentPath.current.Move(dir)
			if p.Value(ng.Get()) == usedSpace {
				continue
			} else if _, ok := visited[*ng]; ok {
				continue
			}

			ngPath := locationPath{
				locations: make([]*helper.Location, len(currentPath.locations)),
				current:   ng,
			}
			copy(ngPath.locations, currentPath.locations)

			queue.PushBack(ngPath)
		}
	}

	// fill the path for puzzle
	p.path = path.locations

	return found
}

func (p *Puzzle) String() string {
	line := "\n🧩\n"

	line += fmt.Sprintf("--- row * col --- %d,%d\n", p.row, p.col)

	for _, row := range p.grid {
		for _, col := range row {
			line += string(col)
		}

		line += "\n"
	}

	line += "🤖" + p.start.String() + "\n"
	line += "⚑" + p.end.String() + "\n"

	return line
}
