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

func (l *Location) Compare(other *Location) int {
	if l.x < other.x {
		return -1
	}

	if l.x > other.x {
		return 1
	}

	if l.y < other.y {
		return -1
	}

	if l.y > other.y {
		return 1
	}

	// equal
	return 0
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
