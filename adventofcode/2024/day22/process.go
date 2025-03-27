package day22

import (
	"fmt"
	"slices"
)

func next(n int) int {
	tmp := n * 64
	n ^= tmp
	n %= 16777216

	tmp = n / 32
	n ^= tmp
	n %= 16777216

	tmp = n * 2048
	n ^= tmp
	n %= 16777216

	return n
}

func secretNumberN(number int, N int) int {
	for range N {
		number = next(number)
	}
	return number
}

func part1(data []int) int {
	solution := 0
	for _, n := range data {
		secret := secretNumberN(n, 2000)
		solution += secret
	}

	return solution
}

func hash(pattern []int) string {
	return fmt.Sprintf("%d,%d,%d,%d", pattern[0], pattern[1], pattern[2], pattern[3])
}

func part2(data []int) int {
	var (
		buyPatterns    = make([]map[string]int, 0)
		uniquePatterns = make([]string, 0)
	)

	for _, n := range data {
		var (
			numPatterns = make(map[string]int)
			buffer      = make([]int, 0)
			prev        = n
		)

		for range 2000 {
			oncePrev := prev % 10
			n := next(prev)
			onceN := n % 10

			diff := onceN - oncePrev
			buffer = append(buffer, diff)

			if len(buffer) > 3 {
				h := hash(buffer)
				if _, ok := numPatterns[h]; !ok {
					numPatterns[h] = onceN
					uniquePatterns = append(uniquePatterns, h)
				}

				buffer = buffer[1:]
			}

			prev = n
		}

		buyPatterns = append(buyPatterns, numPatterns)
	}

	slices.Sort(uniquePatterns)
	uniquePatterns = slices.Compact(uniquePatterns)

	solution := 0
	for _, pattern := range uniquePatterns {
		score := 0
		for _, buyPattern := range buyPatterns {
			score += buyPattern[pattern]
		}

		if score > solution {
			solution = score
		}
	}

	return solution
}
