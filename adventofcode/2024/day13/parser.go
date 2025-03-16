package day13

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	numRgx = regexp.MustCompile(`[0-9]+`)
)

func ParseEquations(data string) []Equation {
	var (
		equations = make([]Equation, 0)
		equation  = Equation{}
	)

	lines := strings.Split(data, "\n")

	for i, line := range lines {
		switch i % 4 {
		case 0: // A data
			nums := numRgx.FindAllStringIndex(line, -1)
			equation.Ax, _ = strconv.Atoi(line[nums[0][0]:nums[0][1]])
			equation.Ay, _ = strconv.Atoi(line[nums[1][0]:nums[1][1]])
		case 1: // B data
			nums := numRgx.FindAllStringIndex(line, -1)
			equation.Bx, _ = strconv.Atoi(line[nums[0][0]:nums[0][1]])
			equation.By, _ = strconv.Atoi(line[nums[1][0]:nums[1][1]])
		case 2: // Prize data
			nums := numRgx.FindAllStringIndex(line, -1)
			equation.X, _ = strconv.Atoi(line[nums[0][0]:nums[0][1]])
			equation.Y, _ = strconv.Atoi(line[nums[1][0]:nums[1][1]])
		case 3: // space
			equations = append(equations, equation)
			equation = Equation{}
		}
	}

	return equations
}
