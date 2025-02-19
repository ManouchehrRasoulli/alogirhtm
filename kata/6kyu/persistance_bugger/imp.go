package kata

func multiply(n int) int {
	x := 1

	for ; n != 0; n /= 10 {
		v := n % 10
		x *= v
	}

	return x
}

func Persistence(n int) int {
	if n < 10 {
		return 0
	}

	return 1 + Persistence(multiply(n))
}
