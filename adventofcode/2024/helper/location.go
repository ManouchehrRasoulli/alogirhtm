package helper

import "fmt"

type Location struct {
	x int
	y int
}

func (l *Location) Get() (int, int) {
	return l.x, l.y
}

func (l *Location) Move(d Direction) {
	l.x += d.x
	l.y += d.y
}

func (l *Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.x, l.y)
}
