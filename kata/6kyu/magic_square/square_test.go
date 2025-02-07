package kata

import "testing"

func Test_MagicSquare(t *testing.T) {
	sq := MagicSquare(5)
	t.Logf("%+v", sq)
}
