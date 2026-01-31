package day10

import (
	"strings"
)

func solveMachineJoltage(m *MachineJoltage) int {
	numCounters := len(m.joltage)
	numButtons := len(m.buttons)

	type rational struct {
		num, den int
	}

	newRat := func(n, d int) rational {
		if d == 0 {
			panic("division by zero")
		}
		if d < 0 {
			n, d = -n, -d
		}
		g := gcd(abs(n), abs(d))
		if g > 0 {
			n, d = n/g, d/g
		}
		return rational{n, d}
	}

	_ = func(a, b rational) rational { // addRat (unused but kept for completeness)
		return newRat(a.num*b.den+b.num*a.den, a.den*b.den)
	}

	subRat := func(a, b rational) rational {
		return newRat(a.num*b.den-b.num*a.den, a.den*b.den)
	}

	mulRat := func(a, b rational) rational {
		return newRat(a.num*b.num, a.den*b.den)
	}

	divRat := func(a, b rational) rational {
		return newRat(a.num*b.den, a.den*b.num)
	}

	matrix := make([][]rational, numCounters)
	for j := 0; j < numCounters; j++ {
		matrix[j] = make([]rational, numButtons+1)
		for i := 0; i < numButtons+1; i++ {
			matrix[j][i] = newRat(0, 1)
		}
		matrix[j][numButtons] = newRat(m.joltage[j], 1) // RHS
	}

	for i, btn := range m.buttons {
		for _, j := range btn {
			matrix[j][i] = newRat(1, 1)
		}
	}

	pivotCol := make([]int, numCounters) // pivotCol[row] = column of pivot for that row, -1 if no pivot
	for i := range pivotCol {
		pivotCol[i] = -1
	}

	row := 0
	for col := 0; col < numButtons && row < numCounters; col++ {
		pivotRow := -1
		for r := row; r < numCounters; r++ {
			if matrix[r][col].num != 0 {
				pivotRow = r
				break
			}
		}
		if pivotRow == -1 {
			continue // No pivot in this column
		}

		matrix[row], matrix[pivotRow] = matrix[pivotRow], matrix[row]
		pivotCol[row] = col

		pivot := matrix[row][col]
		for c := 0; c <= numButtons; c++ {
			matrix[row][c] = divRat(matrix[row][c], pivot)
		}

		for r := 0; r < numCounters; r++ {
			if r != row && matrix[r][col].num != 0 {
				factor := matrix[r][col]
				for c := 0; c <= numButtons; c++ {
					matrix[r][c] = subRat(matrix[r][c], mulRat(factor, matrix[row][c]))
				}
			}
		}
		row++
	}

	isPivotCol := make([]bool, numButtons)
	pivotRows := make(map[int]int) // col -> row mapping for pivot columns
	for r := 0; r < numCounters; r++ {
		if pivotCol[r] >= 0 {
			isPivotCol[pivotCol[r]] = true
			pivotRows[pivotCol[r]] = r
		}
	}

	freeVars := []int{}
	for col := 0; col < numButtons; col++ {
		if !isPivotCol[col] {
			freeVars = append(freeVars, col)
		}
	}

	for r := row; r < numCounters; r++ {
		if matrix[r][numButtons].num != 0 {
			return -1 // No solution
		}
	}

	if len(freeVars) == 0 {
		total := 0
		for col := 0; col < numButtons; col++ {
			if r, ok := pivotRows[col]; ok {
				val := matrix[r][numButtons]
				if val.num < 0 || val.den != 1 || val.num%val.den != 0 {
					return -1 // Non-integer or negative solution
				}
				total += val.num / val.den
			}
		}
		return total
	}

	var bestTotal int = -1

	// Compute a reasonable upper bound for free variables
	// Each button press adds to at least one counter, so max presses is bounded
	maxTarget := 0
	for _, t := range m.joltage {
		if t > maxTarget {
			maxTarget = t
		}
	}

	var search func(freeIdx int, freeVals []int, currentTotal int)
	search = func(freeIdx int, freeVals []int, currentTotal int) {
		if bestTotal >= 0 && currentTotal >= bestTotal {
			return // Prune
		}

		if freeIdx == len(freeVars) {
			valid := true
			total := currentTotal

			for col := 0; col < numButtons && valid; col++ {
				if isPivotCol[col] {
					r := pivotRows[col]
					// x[col] = RHS - sum(coef[free] * freeVal)
					val := matrix[r][numButtons]
					for fi, fv := range freeVars {
						coef := matrix[r][fv]
						val = subRat(val, mulRat(coef, newRat(freeVals[fi], 1)))
					}
					if val.den != 1 || val.num < 0 {
						valid = false
					} else {
						total += val.num
					}
				}
			}

			if valid && (bestTotal < 0 || total < bestTotal) {
				bestTotal = total
			}
			return
		}

		// Use maxTarget as upper bound - no single button needs more presses than the largest target
		for v := 0; v <= maxTarget; v++ {
			freeVals = append(freeVals, v)
			search(freeIdx+1, freeVals, currentTotal+v)
			freeVals = freeVals[:len(freeVals)-1]
		}
	}

	search(0, []int{}, 0)
	return bestTotal
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Part2(input string) int {
	var (
		lines = strings.Split(input, "\n")
		total = 0
	)

	for _, line := range lines {
		m, err := ParseMachineJoltage(line)
		if err != nil {
			panic(err)
		}

		total += solveMachineJoltage(m)
	}

	return total
}
