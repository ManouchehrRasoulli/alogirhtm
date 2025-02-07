package kata

func MagicSquare(n int) [][]int {
	if n%2 == 0 {
		return [][]int{}
	}

	var (
		x      = 0
		y      = (n+1)/2 - 1
		square = make([][]int, n)
	)

	for i := 0; i < n; i++ {
		square[i] = make([]int, n)
	}

	square[x][y] = 1

	for i := 2; i < n*n+1; i++ {
		x_prev := x
		y_prev := y

		if x == 0 {
			x = n - 1
		} else {
			x--
		}

		if y == n-1 {
			y = 0
		} else {
			y++
		}

		for square[x][y] != 0 {
			if x == n-1 {
				x = 0
			} else {
				x = x_prev + 1
			}
			y = y_prev
		}

		square[x][y] = i
	}

	return square
}
