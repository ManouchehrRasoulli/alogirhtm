package kata

import "testing"

func Test_Parse(t *testing.T) {
	v := "iiisdoso"
	t.Log(Parse(v)) // [8, 64]
}
