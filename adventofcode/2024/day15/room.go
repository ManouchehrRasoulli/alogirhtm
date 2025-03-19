package day15

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
)

const (
	wall  = '#'
	box   = 'O'
	bot   = '@'
	empty = '.'
)

var (
	direction = []helper.Direction{
		helper.Top, helper.Right, helper.Bottom, helper.Left,
	}
)

type Room struct {
	ground      [][]rune
	botLocation *helper.Location
	movements   []helper.Direction
}

func (r *Room) Value(x, y int) rune {
	if x < 0 || x >= len(r.ground) || y < 0 || y >= len(r.ground[x]) {
		return wall
	}

	return r.ground[x][y]
}

func (r *Room) SwitchItems(nx, ny, x, y int) {
	nv := r.Value(nx, ny)
	if nv == wall {
		return
	}

	v := r.Value(x, y)
	if v == wall {
		return
	}

	r.ground[nx][ny] = v
	r.ground[x][y] = nv
}

func (r *Room) MoveItemInDirection(x, y int, dir helper.Direction) rune {
	nx, ny := x+dir.X(), y+dir.Y()
	switch r.Value(x, y) {
	case wall:
		return wall
	case empty:
		return empty
	case box:
		next := r.MoveItemInDirection(nx, ny, dir)
		switch next {
		case wall:
			return wall
		case empty:
			r.SwitchItems(nx, ny, x, y)
			return empty
		default:
			return wall
		}
	case bot:
		next := r.MoveItemInDirection(nx, ny, dir)
		switch next {
		case wall:
			return wall
		case box:
			return box
		case empty:
			r.SwitchItems(nx, ny, x, y)
			r.botLocation = helper.NewLocation(nx, ny)
			return empty
		default:
			return wall
		}
	default:
		return wall
	}
}

func (r *Room) MoveBotInDirection(d helper.Direction) {
	x, y := r.botLocation.Get()
	r.MoveItemInDirection(x, y, d)
}

func (r *Room) RunDirections() {
	for _, d := range r.movements {
		r.MoveBotInDirection(d)
	}
}

func (r *Room) GPS() int {
	var (
		gps = 0
	)

	for i, row := range r.ground {
		for j, col := range row {
			switch col {
			case box:
				gps += 100*i + j
			}
		}
	}

	return gps
}

func (r *Room) String() string {
	data := "\n🚹\n"

	for _, row := range r.ground {
		for _, char := range row {
			data += string(char)
		}
		data += "\n"
	}

	if r.botLocation != nil {
		data += fmt.Sprintf("🤖 %s\n", r.botLocation.String())
	}

	for _, row := range r.movements {
		data += string(helper.DirectionToChar[row])
	}

	data += "\n"

	return data
}
