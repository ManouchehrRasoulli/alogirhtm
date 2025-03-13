package day6

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/day4"
	"strings"
)

func ParseRoomWithSimpleGuard(in string) (*Room, *SimpleGuard) {
	var (
		ground = make([][]string, 0)
		lines  = strings.Split(in, "\n")
		guard  = SimpleGuard{}
	)

	for i, line := range lines {
		items := strings.Split(line, "")
		ground = append(ground, items)

		for j, item := range items {
			switch item {
			case "^", ">", "<", "v":
				guard = SimpleGuard{
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

func ParseRoomWithCircuitGuard(in string) (*Room, *CircuitGuard) {
	var (
		ground = make([][]string, 0)
		lines  = strings.Split(in, "\n")
		guard  = CircuitGuard{}
	)

	for i, line := range lines {
		items := strings.Split(line, "")
		ground = append(ground, items)

		for j, item := range items {
			switch item {
			case "^", ">", "<", "v":
				guard = CircuitGuard{
					id:                "K",
					x:                 i,
					y:                 j,
					direction:         day4.CharToDirection[item],
					locationDirection: make(map[string]struct{}),

					startX:         i,
					startY:         j,
					startDirection: day4.CharToDirection[item],
				}
			}
		}
	}

	return &Room{ground: ground}, &guard
}

func LocationToString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func LocationDirectionToString(x, y int, direction day4.Direction) string {
	return fmt.Sprintf("(%d,%d) -> %s", x, y, day4.DirectionToChar[direction])
}
