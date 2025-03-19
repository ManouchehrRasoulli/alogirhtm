package day15

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
)

const (
	boxStart = '['
	boxEnd   = ']'
)

type ScaledRoom struct {
	ground      [][]rune
	botLocation *helper.Location
	movements   []helper.Direction
}

func NewScaledRoom(room *Room) *ScaledRoom {
	s := ScaledRoom{
		ground:      make([][]rune, 0),
		botLocation: nil,
		movements:   room.movements,
	}

	for i, row := range room.ground {
		line := make([]rune, 0)
		for j, col := range row {
			switch col {
			case wall:
				line = append(line, wall, wall)
			case box:
				line = append(line, boxStart, boxEnd)
			case bot:
				s.botLocation = helper.NewLocation(i, 2*j)
				line = append(line, bot, empty)
			case empty:
				line = append(line, empty, empty)
			}
		}
		s.ground = append(s.ground, line)
	}

	return &s
}

func (s *ScaledRoom) String() string {
	data := "\n🚹\n"

	for _, row := range s.ground {
		for _, char := range row {
			data += string(char)
		}
		data += "\n"
	}

	if s.botLocation != nil {
		data += fmt.Sprintf("🤖 %s\n", s.botLocation.String())
	}

	for _, row := range s.movements {
		data += string(helper.DirectionToChar[row])
	}

	data += "\n"

	return data
}

func (s *ScaledRoom) CanMoveItemsVertically(x, y int, dir helper.Direction) bool {
	if dir != helper.Top && dir != helper.Bottom {
		return false
	}

	nx, ny := x+dir.X(), y+dir.Y()
	switch s.Value(x, y) {
	case wall:
		return false
	case empty:
		return true
	case boxStart:
		// have to move both parts
		leftNv := s.CanMoveItemsVertically(nx, ny, dir)
		rightNv := s.CanMoveItemsVertically(nx, ny+1, dir)
		return leftNv && rightNv
	case boxEnd:
		// have to move both parts
		leftNv := s.CanMoveItemsVertically(nx, ny-1, dir)
		rightNv := s.CanMoveItemsVertically(nx, ny, dir)
		return leftNv && rightNv
	case bot:
		return s.CanMoveItemsVertically(nx, ny, dir)
	default:
		return false
	}
}

func (s *ScaledRoom) MoveItemVertically(x, y int, dir helper.Direction) rune {
	if !s.CanMoveItemsVertically(x, y, dir) {
		return wall
	}

	if dir != helper.Top && dir != helper.Bottom {
		return wall
	}

	nx, ny := x+dir.X(), y+dir.Y()
	switch s.Value(x, y) {
	case wall:
		return wall
	case empty:
		return empty
	case boxStart:
		// have to move both parts
		leftNv := s.MoveItemVertically(nx, ny, dir)
		rightNv := s.MoveItemVertically(nx, ny+1, dir)
		if leftNv == rightNv && rightNv == empty {
			s.SwitchItems(nx, ny, x, y)
			s.SwitchItems(nx, ny+1, x, y+1)
			return empty
		} else {
			return wall
		}
	case boxEnd:
		// have to move both parts
		leftNv := s.MoveItemVertically(nx, ny-1, dir)
		rightNv := s.MoveItemVertically(nx, ny, dir)
		if leftNv == rightNv && rightNv == empty {
			s.SwitchItems(nx, ny-1, x, y-1)
			s.SwitchItems(nx, ny, x, y)
			return empty
		} else {
			return wall
		}
	case bot:
		nv := s.MoveItemVertically(nx, ny, dir)
		switch nv {
		case empty:
			s.SwitchItems(nx, ny, x, y)
			s.botLocation = helper.NewLocation(nx, ny)
			return empty
		default:
			return nv
		}
	default:
		return wall
	}
}

func (s *ScaledRoom) MoveItemHorizontally(x, y int, dir helper.Direction) rune {
	if dir != helper.Left && dir != helper.Right {
		return wall
	}

	nx, ny := x+dir.X(), y+dir.Y()
	switch s.Value(x, y) {
	case wall:
		return wall
	case empty:
		return empty
	case boxStart, boxEnd:
		nv := s.MoveItemHorizontally(nx, ny, dir)
		switch nv {
		case empty:
			s.SwitchItems(nx, ny, x, y)
			return empty
		default:
			return nv
		}
	case bot:
		nv := s.MoveItemHorizontally(nx, ny, dir)
		switch nv {
		case empty:
			s.SwitchItems(nx, ny, x, y)
			s.botLocation = helper.NewLocation(nx, ny)
			return empty
		default:
			return nv
		}
	default:
		return wall
	}
}

func (s *ScaledRoom) MoveBotInDirection(d helper.Direction) {
	x, y := s.botLocation.Get()
	switch d {
	case helper.Left, helper.Right:
		s.MoveItemHorizontally(x, y, d)
	case helper.Top, helper.Bottom:
		s.MoveItemVertically(x, y, d)
	}
}

func (s *ScaledRoom) RunDirections() {
	for _, d := range s.movements {
		s.MoveBotInDirection(d)
	}
}

func (s *ScaledRoom) Value(x, y int) rune {
	if x < 0 || x >= len(s.ground) || y < 0 || y >= len(s.ground[x]) {
		return wall
	}

	return s.ground[x][y]
}

func (s *ScaledRoom) SwitchItems(nx, ny, x, y int) {
	nv := s.Value(nx, ny)
	if nv == wall {
		return
	}

	v := s.Value(x, y)
	if v == wall {
		return
	}

	s.ground[nx][ny] = v
	s.ground[x][y] = nv
}

func (s *ScaledRoom) GPS() int {
	var (
		gps = 0
	)

	for i, row := range s.ground {
		for j, col := range row {
			switch col {
			case boxEnd:
				gps += 100*i + j - 1
			}
		}
	}

	return gps
}
