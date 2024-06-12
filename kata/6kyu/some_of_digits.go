package _kyu

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	var sum int
	for x, y := n%10, n/10; y != 0 || x != 0; {
		sum += x
		n = y
		x, y = n%10, n/10
	}

	return DigitalRoot(sum)
}
