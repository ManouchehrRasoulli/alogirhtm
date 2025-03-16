package day13

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestEquation_Solve(t *testing.T) {
	equation := Equation{
		Ax: 94,
		Ay: 34,

		Bx: 22,
		By: 67,

		X: 8400,
		Y: 5400,
	}

	t.Log(equation.Solve()) // 80, 40

	equation = Equation{
		Ax: 26,
		Ay: 66,

		Bx: 67,
		By: 21,

		X: 12748,
		Y: 12176,
	}

	t.Log(equation.Solve()) // -1, -1

	equation = Equation{
		Ax: 17,
		Ay: 86,

		Bx: 84,
		By: 37,

		X: 7870,
		Y: 6450,
	}

	t.Log(equation.Solve()) // 38, 86
}

func TestEquationTestPart1(t *testing.T) {
	data, err := helper.ReadAll("input.test1")
	if err != nil {
		t.Fatal(err)
	}

	equations := ParseEquations(data)
	tokens := 0

	for _, eq := range equations {
		t.Log(eq)
		a, b := eq.Solve()
		if a > 0 && b > 0 {
			tokens += a*3 + b
		}
	}

	t.Log(tokens) // 480
}

func TestEquationPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle")
	if err != nil {
		t.Fatal(err)
	}

	equations := ParseEquations(data)
	tokens := 0

	for _, eq := range equations {
		t.Log(eq)
		a, b := eq.Solve()
		if a > 0 && b > 0 {
			tokens += a*3 + b
		}
	}

	t.Log(tokens) // 37686
}
