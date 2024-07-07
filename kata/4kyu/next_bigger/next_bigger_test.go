package next_bigger

import "testing"

func TestNextBigger(t *testing.T) {
	t.Log(NextBigger(1288))
	t.Log(NextBigger(9))
	t.Log(NextBigger(59884848459853))
	t.Log(NextBigger(1234567890))
}
