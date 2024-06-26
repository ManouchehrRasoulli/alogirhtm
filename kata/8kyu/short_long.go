package _kyu

func Solution(a, b string) string {
	if x, y := len(a), len(b); x > y {
		return b + a + b
	}
	return a + b + a
}
