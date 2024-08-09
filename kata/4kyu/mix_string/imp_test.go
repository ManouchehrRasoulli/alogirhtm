package kata

import "testing"

func TestMix(t *testing.T) {
	t.Log(Mix("aaa bbbbbbb x", "aaaaaaa bbbbbbb nnnn"))
	t.Log(Mix("Are they here", "yes, they are here"))
}
