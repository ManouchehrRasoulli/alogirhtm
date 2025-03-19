package helper

import "fmt"

type Location struct {
	x int
	y int
}

func NewLocation(x int, y int) *Location {
	return &Location{
		x: x,
		y: y,
	}
}

func (l *Location) Get() (int, int) {
	return l.x, l.y
}

func (l *Location) Move(d Direction) *Location {
	return &Location{
		x: l.x + d.x,
		y: l.y + d.y,
	}
}

func (l *Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.x, l.y)
}
