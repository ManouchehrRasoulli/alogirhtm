package day14

func Part1(ground ground) int {
	for TillNorth(ground) {
		continue
	}

	return CalculateWeight(ground)
}
