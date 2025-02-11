package kata

import "fmt"

func permEq(x, y int) bool {
	a, b := []byte(fmt.Sprintf("%d", x)), []byte(fmt.Sprintf("%d", y))

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				a[i], b[j] = '-', '-'
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func SearchPermMult(nMax, k int) int {
	count := 0

	for x := 1; x < nMax; x++ {
		y := x * k
		if y > nMax {
			return count
		}

		if permEq(x, y) {
			fmt.Println("found ", x, " in ", y)
			count++
		}
	}

	return count
}
