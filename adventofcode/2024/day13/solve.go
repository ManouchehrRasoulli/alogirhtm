package day13

import (
	"fmt"
	"math"
)

type Equation struct {
	Ax int
	Ay int

	Bx int
	By int

	X int
	Y int
}

func (e Equation) String() string {
	a, b := e.Solve()
	return fmt.Sprintf("A: X+%d, Y+%d\nB: X+%d, Y+%d\nX=%d, Y=%d\nA :: %d, B :: %d", e.Ax, e.Ay, e.Bx, e.By, e.X, e.Y, a, b)
}

func (e Equation) SolvePart2() (int, int) {
	e.X += 10000000000000
	e.Y += 10000000000000

	a := (e.X*e.By - e.Y*e.Bx) / (e.Ax*e.By - e.Ay*e.Bx) // A = (p_x * b_y - p_y * b_x) / (a_x * b_y - a_y * b_x)
	b := (e.Ax*e.Y - e.Ay*e.X) / (e.Ax*e.By - e.Ay*e.Bx) // B = (a_x * p_y - a_y * p_x) / (a_x * b_y - a_y * b_x)

	if a*e.Ax+b*e.Bx != e.X || a*e.Ay+b*e.By != e.Y {
		a, b = -1, -1
	}

	return a, b
}

func (e Equation) Solve() (int, int) {
	a := (e.X*e.By - e.Y*e.Bx) / (e.Ax*e.By - e.Ay*e.Bx) // A = (p_x * b_y - p_y * b_x) / (a_x * b_y - a_y * b_x)
	b := (e.Ax*e.Y - e.Ay*e.X) / (e.Ax*e.By - e.Ay*e.Bx) // B = (a_x * p_y - a_y * p_x) / (a_x * b_y - a_y * b_x)

	if a*e.Ax+b*e.Bx != e.X || a*e.Ay+b*e.By != e.Y {
		a, b = -1, -1
	}

	return a, b
}

func IsInt(a float64) bool {
	epsilon := 1e-9 // Margin of error
	_, frac := math.Modf(math.Abs(a))
	return frac < epsilon || frac > 1.0-epsilon
}
