package _kyu

func Fortune(f0 int, p float64, c0 int, n int, i float64) bool {
	if c0 > f0 {
		return false
	}

	balance := float64(f0)
	interest := p / 100
	withdraw := float64(c0)
	inflation := i / 100

	for y := 1; y < n; y++ {
		balance = balance + interest*balance - withdraw
		withdraw = withdraw + inflation*withdraw
		if balance < 0 {
			return false
		}
	}

	return true
}
