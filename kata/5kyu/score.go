package _kyu

func Score(dice [5]int) int {
	i := make(map[int]int)

	for _, d := range dice {
		i[d]++
	}

	score := 0
	for k, v := range i {
		if v >= 3 {
			v -= 3
			if k == 1 {
				score += k * 1000
			} else {
				score += k * 100
			}
		}
		i[k] = v
	}

	return score + 100*i[1] + 50*i[5]
}
