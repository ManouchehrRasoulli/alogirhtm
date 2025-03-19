package day6

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
)

type SimpleGuard struct {
	id        string
	x, y      int
	direction helper.Direction
	locations map[string]struct{}
}

func (g *SimpleGuard) String() string {
	return fmt.Sprintf("(%d,%d) -> %s", g.x, g.y, string(helper.DirectionToChar[g.direction]))
}

func (g *SimpleGuard) ChangeToRight() {
	switch g.direction {
	case helper.Top:
		g.direction = helper.Right
	case helper.Right:
		g.direction = helper.Bottom
	case helper.Bottom:
		g.direction = helper.Left
	case helper.Left:
		g.direction = helper.Top
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
		room.Change(g.x, g.y, helper.DirectionToChar[g.direction])
	} else {
		g.ChangeToRight()
		return g.Move(room)
	}

	return true
}
