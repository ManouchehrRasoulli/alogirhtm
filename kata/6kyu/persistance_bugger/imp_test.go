package kata

import "testing"

func Test_multiply(t *testing.T) {
	t.Log(multiply(127)) // 14
	t.Log(multiply(999)) // 729
}

func Test_Persistence(t *testing.T) {
	t.Log(Persistence(999)) // 4
	t.Log(Persistence(4))   // 0
}
