package day3

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"log"
	"strconv"
	"strings"
)

const (
	empty = '.'
)

type Puzzle struct {
	field [][]rune
}

func NewPuzzle(input string) *Puzzle {
	var (
		lines = strings.Split(input, "\n")
		p     = Puzzle{
			field: make([][]rune, 0),
		}
	)

	for _, line := range lines {
		items := make([]rune, 0)
		for _, char := range line {
			items = append(items, char)
		}

		p.field = append(p.field, items)
	}

	return &p
}

func (p *Puzzle) IsAdjacentToSymbol(x, y int) bool {
	for _, d := range helper.AllDirections {
		v := p.Value(x+d.X(), y+d.Y())
		if v == empty {
			continue
		}

		if v >= '0' && v <= '9' {
			continue
		}

		return true
	}

	return false
}

func (p *Puzzle) readLeft(x, y int) []rune {
	v := p.Value(x, y)
	if v == empty || v < '0' || v > '9' {
		return nil
	}

	p.ChangeValue(x, y, empty)

	return append(p.readLeft(x, y-1), v)
}

func (p *Puzzle) readRight(x, y int) []rune {
	v := p.Value(x, y)
	if v == empty || v < '0' || v > '9' {
		return nil
	}

	p.ChangeValue(x, y, empty)

	items := make([]rune, 0)
	items = append(items, v)
	return append(items, p.readRight(x, y+1)...)
}

func (p *Puzzle) NumbersToLeftAndRight(x, y int) int {
	v := p.Value(x, y)
	if v < '0' || v > '9' {
		return 0
	}

	left, right := p.readLeft(x, y-1), p.readRight(x, y+1)

	left = append(left, v)
	left = append(left, right...)

	p.ChangeValue(x, y, empty)

	vi, err := strconv.Atoi(string(left))
	if err != nil {
		log.Fatal(err)
	}

	return vi
}

func (p *Puzzle) ChangeValue(x, y int, v rune) {
	value := p.Value(x, y)
	if value == empty {
		return
	}

	p.field[x][y] = v
}

func (p *Puzzle) Value(x, y int) rune {
	if x < 0 || y < 0 {
		return empty
	}

	if x >= len(p.field) {
		return empty
	}

	if y >= len(p.field[x]) {
		return empty
	}

	return p.field[x][y]
}

func (p *Puzzle) String() string {
	var (
		data = "\n--puzzle--\n"
	)

	for _, line := range p.field {
		data += string(line) + "\n"
	}

	return data
}

func (p *Puzzle) Part1() int {
	var (
		sum = 0
	)

	for x, row := range p.field {
		for y, col := range row {
			if col < '0' || col > '9' {
				continue
			}

			if p.IsAdjacentToSymbol(x, y) {
				sum += p.NumbersToLeftAndRight(x, y)
			}
		}
	}

	return sum
}
