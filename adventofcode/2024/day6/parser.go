package day6

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/day4"
	"strings"
)

type Room struct {
	ground [][]string
}

func (r *Room) WithinBound(x, y int) bool {
	if x < 0 || y < 0 || x >= len(r.ground) || y >= len(r.ground[x]) {
		return false
	}

	return true
}

func (r *Room) Change(x, y int, char string) {
	r.ground[x][y] = char
}

func (r *Room) CanIMoveTo(x, y int) bool {
	if r.ground[x][y] == "#" {
		return false
	}

	return true
}

func (r *Room) String() string {
	var (
		s string
	)

	s += "\n"

	for _, row := range r.ground {
		s += strings.Join(row, "")
		s += "\n"
	}

	return s
}

type Guard struct {
	id        string
	x, y      int
	direction day4.Direction
	locations map[string]struct{}
}

func (g *Guard) String() string {
	return fmt.Sprintf("(%d,%d) -> %s", g.x, g.y, day4.DirectionToChar[g.direction])
}

func (g *Guard) ChangeToRight() {
	switch g.direction {
	case day4.Top:
		g.direction = day4.Right
	case day4.Right:
		g.direction = day4.Bottom
	case day4.Bottom:
		g.direction = day4.Left
	case day4.Left:
		g.direction = day4.Top
	}
}

func (g *Guard) DistinctLocations() map[string]struct{} {
	return g.locations
}

func (g *Guard) Move(room *Room) bool {
	g.locations[fmt.Sprintf("(%d,%d)", g.x, g.y)] = struct{}{}
	
	dx, dy := g.direction.Get()

	nextX := g.x + dx
	nextY := g.y + dy

	if !room.WithinBound(nextX, nextY) {
		return false
	}

	if room.CanIMoveTo(nextX, nextY) {
		g.x, g.y = nextX, nextY
		room.Change(g.x, g.y, day4.DirectionToChar[g.direction])
	} else {
		g.ChangeToRight()
		return g.Move(room)
	}

	return true
}

func ParseRoom(in string) (*Room, *Guard) {
	var (
		ground = make([][]string, 0)
		lines  = strings.Split(in, "\n")
		guard  = Guard{}
	)

	for i, line := range lines {
		items := strings.Split(line, "")
		ground = append(ground, items)

		for j, item := range items {
			switch item {
			case "^", ">", "<", "v":
				guard = Guard{
					id:        "K",
					x:         i,
					y:         j,
					direction: day4.CharToDirection[item],
					locations: make(map[string]struct{}),
				}
			}
		}
	}

	return &Room{ground: ground}, &guard
}
