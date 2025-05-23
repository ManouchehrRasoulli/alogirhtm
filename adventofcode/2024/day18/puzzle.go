package day18

import (
	"container/list"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
)

const (
	freeSpace  = '.'
	usedSpace  = '#'
	trialSpace = 'O' // path which we found
)

type locationPath struct {
	locations []*helper.Location
	current   *helper.Location
}

type Puzzle struct {
	grid [][]rune
	row  int
	col  int
}

func (p *Puzzle) Value(x, y int) rune {
	if x < 0 || y < 0 || x >= len(p.grid) || y >= len(p.grid[x]) {
		return usedSpace
	}

	return p.grid[x][y]
}

func (p *Puzzle) WriteValue(x, y int, value rune) {
	if p.Value(x, y) == usedSpace {
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

func (p *Puzzle) PathCount() int {
	var (
		count int = -1
	)

	for i, row := range p.grid {
		for j, _ := range row {
			if p.Value(i, j) == trialSpace {
				count++
			}
		}
	}

	return count
}

func (p *Puzzle) Bfs() bool {
	var (
		visited = make(map[helper.Location]struct{})
		queue   = list.New()
		end     = helper.NewLocation(p.row, p.col)
		found   = false
	)

	queue.PushBack(helper.NewLocation(0, 0))
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(*helper.Location)

		if _, ok := visited[*current]; ok {
			continue
		}
		visited[*current] = struct{}{}

		if current.Compare(end) == 0 {
			found = true
			break
		}

		if p.Value(current.Get()) == usedSpace {
			continue
		}

		for _, dir := range helper.PathDirections {
			ng := current.Move(dir)
			if p.Value(ng.Get()) == usedSpace {
				continue
			} else if _, ok := visited[*ng]; ok {
				continue
			}

			queue.PushBack(ng)
		}
	}

	return found
}

func (p *Puzzle) BfsWithPath() bool {
	var (
		visited = make(map[helper.Location]struct{})
		queue   = list.New()
		end     = helper.NewLocation(p.row, p.col)
		path    = locationPath{
			locations: make([]*helper.Location, 0),
			current:   helper.NewLocation(0, 0),
		}
		found = false
	)

	queue.PushBack(path)
	for queue.Len() > 0 {
		currentPath := queue.Remove(queue.Front()).(locationPath)
		currentPath.locations = append(currentPath.locations, currentPath.current)

		if _, ok := visited[*currentPath.current]; ok {
			continue
		}
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

	return found
}

func (p *Puzzle) String() string {
	line := "\n🧩\n"

	for _, row := range p.grid {
		for _, col := range row {
			line += string(col)
		}

		line += "\n"
	}

	return line
}
