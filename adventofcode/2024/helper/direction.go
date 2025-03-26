package helper

type Direction struct {
	x int
	y int
}

func (d Direction) Get() (int, int) {
	return d.x, d.y
}

func (d Direction) X() int {
	return d.x
}

func (d Direction) Y() int {
	return d.y
}

var (
	CharToDirection = map[rune]Direction{
		'^': Top,
		'v': Bottom,
		'>': Right,
		'<': Left,
	}
	DirectionToChar = map[Direction]rune{
		Top:    '^',
		Bottom: 'v',
		Right:  '>',
		Left:   '<',
	}
)

var (
	Top = Direction{
		x: -1,
		y: 0,
	}
	Bottom = Direction{
		x: 1,
		y: 0,
	}
	Left = Direction{
		x: 0,
		y: -1,
	}
	Right = Direction{
		x: 0,
		y: 1,
	}
	TopLeft = Direction{
		x: -1,
		y: -1,
	}
	TopRight = Direction{
		x: -1,
		y: 1,
	}
	BottomLeft = Direction{
		x: 1,
		y: -1,
	}
	BottomRight = Direction{
		x: 1,
		y: 1,
	}

	AllDirections  = []Direction{TopLeft, Top, TopRight, Left, Right, BottomLeft, Bottom, BottomRight}
	PathDirections = []Direction{Top, Right, Bottom, Left}
)

func AddUpDirectionToChar(count int, direction Direction) []rune {
	var (
		chars = make([]rune, 0)
	)

	for i := 0; i < count; i++ {
		chars = append(chars, DirectionToChar[direction])
	}

	return chars
}
