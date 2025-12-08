package day1

type Direction string

const (
	Left  Direction = "L"
	Right Direction = "R"
)

type Rotary struct {
	current int
}

func NewRotary(current int) *Rotary {
	return &Rotary{
		current: current,
	}
}

func (r *Rotary) Move(direction Direction, steps int) int {
	switch direction {
	case Left:
		r.current -= steps
	case Right:
		r.current += steps
	}

	r.current %= 100
	return r.current
}
