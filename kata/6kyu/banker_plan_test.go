package _kyu

import "testing"

func TestFortune(t *testing.T) {
	t.Log(Fortune(100_000, 1, 2000, 15, 1))
	t.Log(Fortune(100_000, 1, 10_000, 10, 1))
	t.Log(Fortune(100_000, 1, 9_185, 12, 1))
	t.Log(Fortune(100000000, 1.0, 100000, 50, 1.0))
	t.Log(Fortune(100000000, 1.5, 10000000, 50, 1.0))
	t.Log(Fortune(100000000, 5.0, 1000000, 50, 1.0))
	t.Log(Fortune(9995, 1, 10000, 0, 0))
	t.Log(Fortune(10000, 0, 10000, 2, 0))
}
