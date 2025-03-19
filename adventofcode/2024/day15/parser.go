package day15

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
)

func ParseRoom(data string) *Room {
	r := Room{
		ground:      make([][]rune, 0),
		botLocation: nil,
		movements:   make([]helper.Direction, 0),
	}

	lines := strings.Split(data, "\n")
	for i, line := range lines {
		switch strings.Contains(line, "#") {
		case true: // map
			lineData := make([]rune, 0)
			for j, char := range line {
				lineData = append(lineData, char)
				if char == bot {
					r.botLocation = helper.NewLocation(i, j)
				}
			}
			r.ground = append(r.ground, lineData)
		case false: // command
			for _, char := range line {
				r.movements = append(r.movements, helper.CharToDirection[char])
			}
		}
	}

	return &r
}
