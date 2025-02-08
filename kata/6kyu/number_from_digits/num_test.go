package kata

import "testing"

func TestProcArr(t *testing.T) {
	t.Log(ProcArr([]rune{'1', '2', '2', '3', '2', '3'}))           // [60, 122233, 332221]
	t.Log(ProcArr([]rune{'1', '2', '3', '0', '5', '1', '1', '3'})) // [3360, 1112335, 53321110]
}
