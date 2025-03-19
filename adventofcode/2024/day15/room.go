package day15

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
)

const (
	wall = '#'
	box  = 'O'
	bot  = '@'
)

var (
	direction = []helper.Direction{
		helper.Top, helper.Right, helper.Bottom, helper.Left,
	}
)

type Room struct {
	ground [][]rune
}

func (r *Room) String() string {
	data := "\n"

	for _, row := range r.ground {
		for _, char := range row {
			data += string(char)
		}
		data += "\n"
	}

	data += "\n"

	return data
}
