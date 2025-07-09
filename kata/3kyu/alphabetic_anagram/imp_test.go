package kata

import "testing"

func TestListPosition(t *testing.T) {
	t.Log(ListPosition("ABAB"))       // 2
	t.Log(ListPosition("AAAB"))       // 1
	t.Log(ListPosition("BAAA"))       // 4
	t.Log(ListPosition("QUESTION"))   // 24572
	t.Log(ListPosition("BOOKKEEPER")) // 10743
}
