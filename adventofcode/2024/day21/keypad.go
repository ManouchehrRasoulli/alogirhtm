package day21

import "github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"

const (
	invalidLocation = '!'
)

var (
	pad = [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{' ', '0', 'A'},
	}

	start = helper.NewLocation(3, 2) // `A` pointing at that button
)

type Keypad struct {
	sequence string
	start    *helper.Location
}

func NewKeypad(sequence string) *Keypad {
	return &Keypad{
		sequence: sequence,
		start:    start,
	}
}

func (k *Keypad) ValueLocation(value rune) *helper.Location {
	panic("not implemented")
}

func (k *Keypad) Value(location *helper.Location) rune {
	x, y := location.Get()
	if x < 0 || y < 0 || x > 3 || y > 3 {
		return invalidLocation
	}

	if x == 2 && y == 0 {
		return invalidLocation
	}

	return pad[y][x]
}

func (k *Keypad) PathFromAtoB(a *helper.Location, b *helper.Location) []rune {
	// this sequence will find the path from position a to b
	var (
		ax, ay              = a.Get()
		bx, by              = b.Get()
		upDirectionCount    = 0
		rightDirectionCount = 0
		downDirectionCount  = 0
		leftDirectionCount  = 0
	)

	if ax != bx {
		if ax < bx {
			downDirectionCount += bx - ax
		} else {
			upDirectionCount += ax - bx
		}
	}

	if ay != by {
		if ay < by {
			rightDirectionCount += by - ay
		} else {
			leftDirectionCount += ay - by
		}
	}

	panic("not implemented yet !")
}
