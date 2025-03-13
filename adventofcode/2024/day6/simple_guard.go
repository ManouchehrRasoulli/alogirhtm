package day6

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/day4"
)

type SimpleGuard struct {
	id        string
	x, y      int
	direction day4.Direction
	locations map[string]struct{}
}

func (g *SimpleGuard) String() string {
	return fmt.Sprintf("(%d,%d) -> %s", g.x, g.y, day4.DirectionToChar[g.direction])
}

func (g *SimpleGuard) ChangeToRight() {
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

func (g *SimpleGuard) DistinctLocations() map[string]struct{} {
	return g.locations
}

func (g *SimpleGuard) Move(room *Room) bool {
	if len(g.locations) == 0 {
		g.locations[LocationToString(g.x, g.y)] = struct{}{} // starting location
	}

	dx, dy := g.direction.Get()

	nextX := g.x + dx
	nextY := g.y + dy

	if !room.WithinBound(nextX, nextY) {
		g.locations[LocationToString(g.x, g.y)] = struct{}{} // final location
		return false
	}

	if room.CanIMoveTo(nextX, nextY) {
		g.locations[LocationToString(g.x, g.y)] = struct{}{}
		g.x, g.y = nextX, nextY
		room.Change(g.x, g.y, day4.DirectionToChar[g.direction])
	} else {
		g.ChangeToRight()
		return g.Move(room)
	}

	return true
}
