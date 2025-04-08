package _00_arranged

import "log"

func FindBlueCount() int {
	b, n := 15, 21

	for n < 1_000_000_000_000 {
		//  1_070_379_110_497
		b, n = 3*b+2*n-2, 4*b+3*n-3
	}

	log.Println(b, n)

	return b
}
