package change

import "testing"

func TestChange(t *testing.T) {
	coins := []int{25, 10, 5, 1}
	t.Log(change(coins, 30)) // 2
}
