package day14

const totalCycles = 1_000_000_000

func Part2(ground ground) int {
	seen := make(map[string]int)

	for i := 0; i < totalCycles; i++ {
		key := ground.String()
		if start, ok := seen[key]; ok {
			cycleLength := i - start
			remaining := (totalCycles - i) % cycleLength
			for r := 0; r < remaining; r++ {
				Cycle(ground)
			}
			return CalculateWeight(ground)
		}
		seen[key] = i

		Cycle(ground)
	}

	return CalculateWeight(ground)
}
