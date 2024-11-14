package kata

import "testing"

func Test_SumOfSquares(t *testing.T) {
	t.Log(SumOfSquares(12)) // 3
	t.Log(SumOfSquares(13)) // 2
	t.Log(SumOfSquares(18)) // 2
	t.Log(SumOfSquares(19)) // 3
	t.Log(SumOfSquares(40)) // 2
	t.Log(SumOfSquares(661915703))
}
