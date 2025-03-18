package day14

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	numberRegex = regexp.MustCompile(`[-]{0,1}[0-9]+`)
)

type Bot struct {
	X, Y                 int
	XVelocity, YVelocity int
}

func NewBotFromString(s string) (*Bot, error) {
	// sample input :: p=0,4 v=3,-3
	var (
		rawData = numberRegex.FindAllString(s, -1)
	)

	b := Bot{
		X:         0,
		Y:         0,
		XVelocity: 0,
		YVelocity: 0,
	}

	for i, v := range rawData {
		vn, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		switch i % 4 {
		case 0:
			b.Y = vn
			break
		case 1:
			b.X = vn
			break
		case 2:
			b.YVelocity = vn
			break
		case 3:
			b.XVelocity = vn
			break
		}
	}

	return &b, nil
}

func (b *Bot) MoveInGrid(x, y int) {
	b.X += b.XVelocity
	b.Y += b.YVelocity

	if b.X < 0 {
		b.X = x + b.X
	}

	if b.Y < 0 {
		b.Y = y + b.Y
	}

	if b.X >= x {
		b.X %= x
	}

	if b.Y >= y {
		b.Y %= y
	}
}

func (b *Bot) String() string {
	return fmt.Sprintf("🤖 (%d, %d) -> (%d, %d)", b.X, b.Y, b.XVelocity, b.YVelocity)
}

func CountQuadrants(robots []*Bot, x, y int) (int, int, int, int) {
	qtl := 0
	qtr := 0
	qbl := 0
	qbr := 0

	for _, r := range robots {
		if r.X < x/2 && r.Y < y/2 {
			qtl++
		} else if r.X < x/2 && r.Y >= y-y/2 {
			qtr++
		} else if r.X >= x-x/2 && r.Y < y/2 {
			qbl++
		} else if r.X >= x-x/2 && r.Y >= y-y/2 {
			qbr++
		}
	}

	return qtl, qtr, qbl, qbr
}

type Field struct {
	matrix [][]int
	x, y   int
}

func NewField(x, y int) *Field {
	f := Field{
		matrix: make([][]int, 0),
		x:      x,
		y:      y,
	}

	for i := 0; i < x; i++ {
		data := make([]int, 0)
		for j := 0; j < y; j++ {
			data = append(data, 0)
		}

		f.matrix = append(f.matrix, data)
	}

	return &f
}

func (f *Field) Reset() {
	for i := 0; i < f.x; i++ {
		for j := 0; j < f.y; j++ {
			f.matrix[i][j] = 0
		}
	}
}

func (f *Field) DrawBots(bots []*Bot) {
	f.Reset()

	for _, bot := range bots {
		f.matrix[bot.X][bot.Y]++
	}
}

func (f *Field) String() string {
	data := "\n"

	for _, row := range f.matrix {
		for _, cell := range row {
			if cell > 0 {
				// data += fmt.Sprintf("%d", cell)
				data += "*"
			} else {
				data += "."
			}
		}
		data += "\n"
	}

	data += "\n"

	return data
}
