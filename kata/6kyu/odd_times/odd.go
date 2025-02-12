package kata

func FindOdd(seq []int) int {
	occ := make(map[int]int)

	for _, n := range seq {
		occ[n]++
		if occ[n] > 1 {
			delete(occ, n)
			continue
		}
	}

	for k, _ := range occ {
		return k
	}

	return -1000
}
