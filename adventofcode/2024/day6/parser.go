package day6

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
)

func ParseRoomWithSimpleGuard(in string) (*Room, *SimpleGuard) {
	var (
		ground = make([][]rune, 0)
		lines  = strings.Split(in, "\n")
		guard  = SimpleGuard{}
	)

	for i, line := range lines {
		ground = append(ground, []rune(line))

		for j, item := range line {
			switch item {
			case '^', '>', '<', 'v':
				guard = SimpleGuard{
					id:        "K",
					x:         i,
					y:         j,
					direction: helper.CharToDirection[item],
					locations: make(map[string]struct{}),
				}
			}
		}
	}

	return &Room{ground: ground}, &guard
}

func ParseRoomWithCircuitGuard(in string) (*Room, *CircuitGuard) {
	var (
		ground = make([][]rune, 0)
		lines  = strings.Split(in, "\n")
		guard  = CircuitGuard{}
	)

	for i, line := range lines {
		ground = append(ground, []rune(line))
		for j, item := range line {
			switch item {
			case '^', '>', '<', 'v':
				guard = CircuitGuard{
					id:                "K",
					x:                 i,
					y:                 j,
					direction:         helper.CharToDirection[item],
					locationDirection: make(map[string]struct{}),

					startX:         i,
					startY:         j,
					startDirection: helper.CharToDirection[item],
				}
			}
		}
	}

	return &Room{ground: ground}, &guard
}

func LocationToString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func LocationDirectionToString(x, y int, direction helper.Direction) string {
	return fmt.Sprintf("(%d,%d) -> %s", x, y, string(helper.DirectionToChar[direction]))
}
