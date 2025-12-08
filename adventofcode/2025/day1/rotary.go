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
		if r.current < 0 {
			r.current += 100
		}
	case Right:
		r.current += steps
		if r.current > 99 {
			r.current -= 100
		}
	}
	r.current %= 100

	return r.current
}
