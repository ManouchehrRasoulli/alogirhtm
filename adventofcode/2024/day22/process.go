package day22

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
