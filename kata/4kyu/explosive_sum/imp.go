package kata

import (
	"fmt"
)

var preComputed = make(map[string]uint64)

func ExpSum(n uint64) uint64 {
	fmt.Println(n)
	return explosiveSum(n, n)
}

func explosiveSum(n uint64, max uint64) uint64 {
	if n == 0 {
		return 1
	}

	var sum uint64

	for i := _min(n, max); i > 0; i-- {
		a, b := n-i, i

		ok, val := restore(a, b)
		if ok {
			sum += val
			continue
		}

		val = explosiveSum(a, b)
		sum += val

		store(a, b, val)
	}

	return sum
}

func store(a, b, v uint64) {
	preComputed[fmt.Sprintf("%d-%d", a, b)] = v
}

func restore(a, b uint64) (bool, uint64) {
	ok, val := preComputed[fmt.Sprintf("%d-%d", a, b)]
	return val, ok
}

func _min(a, b uint64) uint64 {
	if a > b {
		return b
	}

	return a
}
