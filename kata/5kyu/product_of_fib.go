package _kyu

import "fmt"

var fibCalculated []uint64

func init() {
	fibCalculated = make([]uint64, 0)
	fibCalculated = append(fibCalculated, 0)
	fibCalculated = append(fibCalculated, 1)
	fibCalculated = append(fibCalculated, 1)
}

func fibX(x int) uint64 {
	if x < len(fibCalculated) { // means we already calculated this item !
		return fibCalculated[x]
	}

	for i := len(fibCalculated); i <= x; i++ {
		current := fibCalculated[i-1] + fibCalculated[i-2]
		fibCalculated = append(fibCalculated, current)
	}

	return fibCalculated[x]
}

func ProductFib(prod uint64) [3]uint64 {
	fmt.Println("--> input: ", prod)
	var (
		x, y        = 0, 1
		p    uint64 = 0
		fibx uint64
		fiby uint64
	)

	for p < prod {
		fibx = fibX(x)
		fiby = fibX(y)
		p = fibx * fiby
		if p == prod {
			return [3]uint64{fibx, fiby, 1}
		}
		x = y
		y++
	}

	return [3]uint64{fibx, fiby, 0}
}
