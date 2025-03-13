package day6

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/day4"
)

type CircuitGuard struct {
	id                string
	x, y              int
	direction         day4.Direction
	locationDirection map[string]struct{} // if i am in location with same direction then i am in a circle !

	startX, startY int
	startDirection day4.Direction
}

func (g *CircuitGuard) Reset() {
	g.x, g.y = g.startX, g.startY
	g.direction = g.startDirection
	g.locationDirection = make(map[string]struct{})
}

func (g *CircuitGuard) String() string {
	return fmt.Sprintf("(%d,%d) -> %s", g.x, g.y, day4.DirectionToChar[g.direction])
}

func (g *CircuitGuard) ChangeToRight() {
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

func (g *CircuitGuard) Move(room *Room) bool {
	dx, dy := g.direction.Get()

	nextX := g.x + dx
	nextY := g.y + dy

	if !room.WithinBound(nextX, nextY) {
		return false
	}

	if room.CanIMoveTo(nextX, nextY) {
		g.x, g.y = nextX, nextY
	} else {
		g.ChangeToRight()
		return g.Move(room)
	}

	return true
}

func (g *CircuitGuard) AmIInACircuit(room *Room) bool {
	defer g.Reset()

	g.locationDirection[LocationDirectionToString(g.x, g.y, g.direction)] = struct{}{}

	for g.Move(room) {
		if _, circle := g.locationDirection[LocationDirectionToString(g.x, g.y, g.direction)]; circle {
			return true
		}

		g.locationDirection[LocationDirectionToString(g.x, g.y, g.direction)] = struct{}{}
	}

	return false
}

func (g *CircuitGuard) CountCircuits(room *Room) int {
	circuits := 0

	iC, jC := room.Size()
	for i := 0; i < iC; i++ {
		for j := 0; j < jC; j++ {
			if i == g.startX && j == g.startY {
				continue
			}

			if !room.CanIMoveTo(i, j) {
				continue
			}

			room.MakeObstacle(i, j)
			if g.AmIInACircuit(room) {
				circuits++
			}
			room.MakeFree(i, j)
		}
	}

	return circuits
}
