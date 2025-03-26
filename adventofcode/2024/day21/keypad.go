package day21

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"log"
)

const (
	invalidLocation = '!'
	noValue         = ' '
	confirm         = 'A'
)

var (
	pad = [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{' ', '0', 'A'},
	}
)

type Keypad struct {
	sequence string
}

func NewKeypad(sequence string) *Keypad {
	sequence = string(confirm) + sequence // start from `A` and finish with `A`
	return &Keypad{
		sequence: sequence,
	}
}

func (k *Keypad) valueLocation(value rune) *helper.Location {
	for i, row := range pad {
		for j, col := range row {
			if col == value {
				return helper.NewLocation(i, j)
			}
		}
	}

	log.Fatal("looking for invalid character at keypad !!", string(value), k.sequence)
	return nil
}

func (k *Keypad) pathFromAtoB(a *helper.Location, b *helper.Location) []rune {
	// this sequence will find the path from position a to b
	var (
		direction           = make([]rune, 0)
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

	direction = append(direction, helper.AddUpDirectionToChar(leftDirectionCount, helper.Left)...)
	direction = append(direction, helper.AddUpDirectionToChar(rightDirectionCount, helper.Right)...)
	direction = append(direction, helper.AddUpDirectionToChar(downDirectionCount, helper.Bottom)...)
	direction = append(direction, helper.AddUpDirectionToChar(upDirectionCount, helper.Top)...)

	return direction
}

func (k *Keypad) CharToDirections() []rune {
	var (
		path = make([]rune, 0)
	)

	log.Println("keypad compute directions ::", k.sequence)

	for i := 1; i < len(k.sequence); i++ {
		ac, bc := k.sequence[i-1], k.sequence[i]
		al, bl := k.valueLocation(rune(ac)), k.valueLocation(rune(bc))
		log.Println(string(ac), al, string(bc), bl)

		direction := k.pathFromAtoB(al, bl)
		path = append(path, direction...)
		path = append(path, confirm)
	}

	return path
}
