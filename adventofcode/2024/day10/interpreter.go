package day10

import (
	"container/list"
	"fmt"
	"log"
)

const (
	noValue  int = -10
	minValue     = 0
	maxValue     = 9
)

type Location struct {
	x int
	y int
}

func (l *Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.x, l.y)
}

type Island struct {
	field      [][]int
	trailheads map[Location]struct{}
}

func (i *Island) Value(location Location) int {
	if location.x < 0 || location.y < 0 || location.x >= len(i.field) || location.y >= len(i.field[location.x]) {
		return -3
	}

	return i.field[location.x][location.y]
}

func (i *Island) Neighbors(location Location) []Location {
	var (
		nextVal   = i.Value(location) + 1 // next value we are looking for
		neighbors []Location
		up        = Location{location.x - 1, location.y}
		down      = Location{location.x + 1, location.y}
		left      = Location{location.x, location.y - 1}
		right     = Location{location.x, location.y + 1}
	)

	if nextVal == noValue+1 {
		return neighbors
	}

	if i.Value(up) == nextVal {
		neighbors = append(neighbors, up)
	}

	if i.Value(down) == nextVal {
		neighbors = append(neighbors, down)
	}

	if i.Value(left) == nextVal {
		neighbors = append(neighbors, left)
	}

	if i.Value(right) == nextVal {
		neighbors = append(neighbors, right)
	}

	return neighbors
}

func (i *Island) TrailheadScoresSumPart1() int {
	var (
		trailheadScore = 0
	)

	log.Println("Trailhead Score ::")
	for trailhead, _ := range i.trailheads {
		score := i.Bfs(trailhead)
		log.Println(trailhead, score)
		trailheadScore += score
	}

	return trailheadScore
}

func (i *Island) Bfs(location Location) int {
	var (
		trailheadScore      = 0
		queue               = list.New()
		locationHealVisited = make(map[Location]struct{})
	)

	queue.PushBack(location)
	for queue.Len() > 0 {
		loc := queue.Remove(queue.Front()).(Location)
		neighbors := i.Neighbors(loc)
		for _, neighbor := range neighbors {
			if i.Value(neighbor) == maxValue {
				if _, ok := locationHealVisited[neighbor]; ok {
					continue
				}
				trailheadScore++
				locationHealVisited[neighbor] = struct{}{}
				continue
			} else if i.Value(neighbor) == noValue {
				continue
			} else if i.Value(neighbor) != i.Value(loc)+1 {
				continue
			}

			queue.PushBack(neighbor)
		}
	}

	return trailheadScore
}

func (i *Island) String() string {
	lines := "\n"

	for _, row := range i.field {
		lines += fmt.Sprintf("%+v\n", row)
	}

	lines += "\nTrailheads:\n"
	for row, _ := range i.trailheads {
		lines += row.String() + " "
	}

	return lines
}
